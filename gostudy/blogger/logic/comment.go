package logic

import (
	"fmt"
	"github.com/gnufree/gostudy/blogger/dal/db"
	"github.com/gnufree/gostudy/blogger/model"
	"time"
)

func InsertComment(comment, username,email string,articleId int64) (err error) {
	// 1. 首先，验证文章id是否合法
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n",err)
		return
	}

	if exist == false {
		err = fmt.Errorf("article id:%d not fount",articleId)
		return
	}
	// 2.调用dal InsertComment 进行评论内容插入
	var c model.Comment
	c.ArticleId = articleId
	c.Content = comment
	c.Username = username
	c.CreateTime = time.Now()
	c.Status = 1
	err = db.InsertComment(&c)


	return
}

func GetCommentList(articleId int64) (commentList []*model.Comment, err error) {
	// 1. 首先，验证文章id是否合法
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n",err)
		return
	}

	if exist == false {
		err = fmt.Errorf("article id:%d not fount",articleId)
		return
	}
	// 2.调用dal GetCommentList 获取评论列表
	commentList, err = db.GetCommentList(articleId,0,100)
	return
}