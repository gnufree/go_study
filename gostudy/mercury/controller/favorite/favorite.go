package favorite

import (
	"strconv"
	"strings"

	"github.com/gnufree/gostudy/mercury/dal/db"

	"github.com/gnufree/gostudy/mercury/middleware/account"

	"github.com/gnufree/gostudy/mercury/id_gen"

	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/common"
	"github.com/gnufree/gostudy/mercury/util"
)

func AddDirHandle(c *gin.Context) {
	var favoriteDir common.FavoriteDir
	err := c.BindJSON(&favoriteDir)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	favoriteDir.DirName = strings.TrimSpace(favoriteDir.DirName)
	if len(favoriteDir.DirName) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("invalid dir name:%v", favoriteDir.DirName)
		return
	}
	logger.Debug("bind json success, favoriteDir:%#v", favoriteDir)
	dir_id, err := id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("id_gen.GetId failed, favoriteDir:%#v,err:%v", favoriteDir, err)
		return
	}
	favoriteDir.DirId = int64(dir_id)

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	favoriteDir.UserId = userId
	err = db.CreateFavoriteDir(&favoriteDir)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("CreateFavoriteDir failed, favoriteDir:%#v,err:%v", favoriteDir, err)
		return
	}
	util.ResponseSuccess(c, nil)
}

func AddFavoriteHandle(c *gin.Context) {

	var favorite common.Favorite
	err := c.BindJSON(&favorite)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	if favorite.DirId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("invalid favorite:%#v", favorite)
		return
	}

	logger.Debug("bind json success, favorite:%#v", favorite)

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	favorite.UserId = userId
	err = db.CreateFavorite(&favorite)
	if err != nil {
		if err == db.ErrRecordExists {
			util.ResponseError(c, util.ErrCodeRecordExist)
		} else {
			util.ResponseError(c, util.ErrCodeServerBusy)
		}
		logger.Error("CreateFavorite failed, favorite:%#v,err:%v", favorite, err)
		return
	}
	util.ResponseSuccess(c, nil)

}
func DirListHandle(c *gin.Context) {
	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	favoriteDirList, err := db.GetFavoriteDirList(userId)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetFavoriteDirList failed, userId:%v, err:%v", userId, err)
		return
	}
	util.ResponseSuccess(c, favoriteDirList)
}
func FavoriteListHandle(c *gin.Context) {

	dirIdStr, ok := c.GetQuery("dir_id")
	dirIdStr = strings.TrimSpace(dirIdStr)
	if ok == false || len(dirIdStr) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid dir id, val:%v", dirIdStr)
		return
	}
	logger.Debug("get query dir_id success, val:%v", dirIdStr)
	dirId, err := strconv.ParseInt(dirIdStr, 10, 64)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid dir id, val:%v", dirIdStr)
		return
	}
	logger.Debug("get query dir_id success, val:%v", dirIdStr)

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

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}
	favoriteList, err := db.GetFavoriteList(userId, dirId, offset, limit)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetFavoriteList failed, dir_id, user_id,err:%v", dirId, userId, err)
		return
	}

	var answerIdList []int64

	for _, v := range favoriteList {
		answerIdList = append(answerIdList, v.AnswerId)
	}
	answerList, err := db.MGetAnswer(answerIdList)
	if err != nil {
		logger.Error("db.MGetAnswer failed, answer_idx:%v, err:%v",
			answerList, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	var userIdList []int64
	for _, v := range answerList {
		userIdList = append(userIdList, v.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("db.GetUserInfoList failed, user_idx:%v, err:%v",
			userIdList, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	apiAnswerList := &common.ApiAnswerList{}
	for _, v := range answerList {
		apiAnswer := &common.ApiAnswer{}
		apiAnswer.Answer = *v

		for _, user := range userInfoList {
			if user.UserId == v.AuthorId {
				apiAnswer.AuthorName = user.Username
				break
			}
		}
		apiAnswerList.AnswerList = append(apiAnswerList.AnswerList, apiAnswer)
	}

	util.ResponseSuccess(c, apiAnswerList)
}
