package category

import (
	"strconv"

	"github.com/gnufree/gostudy/mercury/common"

	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/dal/db"
	"github.com/gnufree/gostudy/mercury/util"
)

func GetCategoryListHandle(c *gin.Context) {

	categoryList, err := db.GetCategoryList()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(c, categoryList)
}

func GetQuestionListHandle(c *gin.Context) {

	categoryIdStr, ok := c.GetQuery("category_id")
	if !ok {
		logger.Error("invalid category_id, not found, category_id:%v", categoryIdStr)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		logger.Error("invalid category_id, strconv.ParseInt failed, err:%v str:%v",
			err, categoryIdStr)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	questionList, err := db.GetQuestionList(categoryId)
	if err != nil {
		logger.Error("get question list failed, cagegory_id:%v, err:%v", categoryId, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	var userIdList []int64
	userIdMap := make(map[int64]bool, 16)
	for _, question := range questionList {
		_, ok := userIdMap[question.AuthorId]
		if ok {
			continue
		}
		userIdMap[question.AuthorId] = true
		userIdList = append(userIdList, question.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("get userInfoList failed,userInfoList:%v, err:%v",
			userIdList, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	var apiQuestionList []*common.ApiQuestion
	for _, question := range questionList {
		var apiQuestion = &common.ApiQuestion{}
		apiQuestion.Question = *question
		apiQuestion.CreateTimeStr = apiQuestion.CreateTime.Format("2006/1/2 15:04:05")

		for _, userInfo := range userInfoList {
			if question.AuthorId == userInfo.UserId {
				apiQuestion.AuthorName = userInfo.Username
				break
			}
		}
		apiQuestionList = append(apiQuestionList, apiQuestion)
	}
	util.ResponseSuccess(c, apiQuestionList)
}
