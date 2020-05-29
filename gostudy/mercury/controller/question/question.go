package question

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/common"
	"github.com/gnufree/gostudy/mercury/dal/db"
	"github.com/gnufree/gostudy/mercury/filter"
	"github.com/gnufree/gostudy/mercury/id_gen"
	"github.com/gnufree/gostudy/mercury/middleware/account"
	"github.com/gnufree/gostudy/mercury/util"
)

func QuestionSubmitHandle(c *gin.Context) {

	var question common.Question

	err := c.BindJSON(&question)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	logger.Debug("bind json success, question:%#v", question)
	result, hit := filter.Replace(question.Caption, "***")
	if hit {
		logger.Error("caption is hit filter,result:%v", result)
		util.ResponseError(c, util.ErrCodeCaptionHit)
		return
	}
	logger.Debug("caption filter success, result:%#v", result)

	result, hit = filter.Replace(question.Content, "***")
	if hit {
		logger.Error("content is hit filter,result:%v", result)

		util.ResponseError(c, util.ErrCodeContentHit)
		return
	}
	logger.Debug("content filter success, result:%#v", result)

	qid, err := id_gen.GetId()
	if err != nil {
		logger.Error("generate question id failed,err:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	question.QuestionId = int64(qid)
	userId, err := account.GetUserId(c)
	if err != nil || userId <= 0 {
		logger.Error("user is not login, err:%v", err)
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	question.AuthorId = userId

	err = db.CreateQuestion(&question)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	logger.Debug("create question success,question:%#v", question)
	util.ResponseSuccess(c, nil)
	util.SendKafka("mercury_topic", question)
}

func QuestionDetailHandle(c *gin.Context) {

	questionIdStr, ok := c.GetQuery("question_id")
	if !ok {
		logger.Error("invalid question_id not fount. question_id:%v", questionIdStr)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	questionId, err := strconv.ParseInt(questionIdStr, 10, 64)
	if err != nil {
		logger.Error("strcont.PartInt failed, err:%v, str:%v", err, questionIdStr)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	question, err := db.GetQuestion(questionId)
	if err != nil {
		logger.Error("get question failed, str:%v ,err:%v", questionId, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	categoryMap, err := db.MGetCategory([]int64{question.CategoryId})
	if err != nil {
		logger.Error("get category failed, err:%v, question:%v", err, question)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	category, ok := categoryMap[question.CategoryId]
	if !ok {
		logger.Error("get category failed, err:%v, question:%v", err, question)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	userInfoList, err := db.GetUserInfoList([]int64{question.AuthorId})
	if err != nil || len(userInfoList) == 0 {
		logger.Error("get user info list failed, user_idx:%#v, err:%v", question.AuthorId, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	apiQuestionDetail := &common.ApiQuestionDetail{}
	apiQuestionDetail.Question = *question
	apiQuestionDetail.AuthorName = userInfoList[0].Username
	apiQuestionDetail.CategoryName = category.CategoryName

	util.ResponseSuccess(c, apiQuestionDetail)
}
