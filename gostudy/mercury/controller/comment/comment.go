package comment

import (
	"html"
	"strconv"
	"strings"

	"github.com/gnufree/gostudy/mercury/dal/db"

	"github.com/gnufree/gostudy/mercury/id_gen"

	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/common"
	"github.com/gnufree/gostudy/mercury/middleware/account"
	"github.com/gnufree/gostudy/mercury/util"
)

const (
	MinCommentContentSize = 10
)

func PostCommentHandle(c *gin.Context) {

	var comment common.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	logger.Debug("bind json success, comment:%#v", comment)
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	userId, err := account.GetUserId(c)
	if err != nil {
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	comment.AuthorId = userId
	// 1、针对content做一个转义，防止xss漏洞
	comment.Content = html.EscapeString(comment.Content)

	// 2、生成评论id
	cid, err := id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("id_gen.GetId failed, comment:%#v, err:%v", comment, err)
		return
	}
	comment.CommentId = int64(cid)
	err = db.CreatePostComment(&comment)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("db.CreatePostComment failed, comment:%#v, err:%v", comment, err)
		return
	}
	util.ResponseSuccess(c, nil)
}

func PostReplyHandle(c *gin.Context) {
	var comment common.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	logger.Debug("bind json success, comment:%#v", comment)
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 ||
		comment.ReplyCommentId == 0 || comment.ParentId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("len(comment.Content):%v, qid:%v, invalid param",
			len(comment.Content), comment.QuestionId)
		return
	}
	userId, err := account.GetUserId(c)
	if err != nil {
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	comment.AuthorId = userId
	// 1、针对content做一个转义，防止xss漏洞
	comment.Content = html.EscapeString(comment.Content)

	// 2、生成评论id
	cid, err := id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("id_gen.GetId failed, comment:%#v, err:%v", comment, err)
		return
	}
	// 3、根据ReplyCommentId, 查询这个ReplyCommentId的author_id, 也就是ReplyAuthorId
	comment.CommentId = int64(cid)
	err = db.CreateReplyComment(&comment)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("db.CreatePostComment failed, comment:%#v, err:%v", comment, err)
		return
	}
	util.ResponseSuccess(c, nil)
}

func CommentListHandle(c *gin.Context) {
	answerIdStr, ok := c.GetQuery("answer_id")
	answerIdStr = strings.TrimSpace(answerIdStr)
	if ok == false || len(answerIdStr) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid answer id, val:%v", answerIdStr)
		return
	}
	logger.Debug("get query answer_id success, val:%v", answerIdStr)
	answerId, err := strconv.ParseInt(answerIdStr, 10, 64)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid answer id, val:%v", answerIdStr)
		return
	}
	logger.Debug("get query answer_id success, val:%v", answerIdStr)

	//解析offset
	var offset int64
	offsetStr, ok := c.GetQuery("offset")
	offsetStr = strings.TrimSpace(offsetStr)
	if ok == false || len(offsetStr) == 0 {
		offset = 0
		logger.Error("invalid offset, val:%v", offsetStr)
	}
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
		logger.Error("invalid offset, val:%v", offsetStr)
	}
	logger.Debug("get query offset success, val:%v", offsetStr)
	// 解析limit
	var limit int64
	limitStr, ok := c.GetQuery("limit")
	limitStr = strings.TrimSpace(limitStr)
	if ok == false || len(limitStr) == 0 {
		offset = 0
		logger.Error("invalid limit, val:%v", limitStr)
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		offset = 0
		logger.Error("invalid limit, val:%v", limitStr)
	}
	logger.Debug("get query limit success, val:%v", limitStr)

	// 获取一级评论列表
	commentList, count, err := db.GetCommentList(answerId, offset, limit)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetCommentList failed, answer_id:%v, err:%v", answerId, err)
		return
	}

	// 获取评论用户信息
	var userIdList []int64
	for _, v := range commentList {
		userIdList = append(userIdList, v.AuthorId, v.ReplyAuthorId)
	}
	userList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetUserInfoList failed, answer_id:%v, err:%v", answerId, err)
		return
	}

	userInfoMap := make(map[int64]*common.UserInfo, len(userIdList))
	for _, user := range userList {
		userInfoMap[user.UserId] = user
	}

	for _, v := range commentList {
		user, ok := userInfoMap[v.AuthorId]
		if ok {
			v.AuthorName = user.Username
		}

		user, ok = userInfoMap[v.ReplyAuthorId]
		if ok {
			v.ReplyAuthorName = user.Username
		}
	}

	var apiCommentList = &common.ApiCommentList{}
	apiCommentList.Count = count
	apiCommentList.CommentList = commentList

	util.ResponseSuccess(c, apiCommentList)
}
