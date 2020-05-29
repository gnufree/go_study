package db

import (
	"testing"
	"time"

	"github.com/gnufree/gostudy/blogger/model"
)

func init() {
	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}

}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Summary = `MySQL是一个关系型数据库管理系统,由瑞典MySQL AB 公司开发
	属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统之一
	在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System
	关系数据库管理系统) 应用软件之一`
	article.ArticleInfo.Title = "Mysql 详情"
	article.ArticleInfo.Username = "智文军"
	article.ArticleInfo.ViewCount = 1
	article.Category.CategoryId = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("insert article failed, err:%v\n", err)
		return
	}
	t.Logf("insert article success, articleId:%d\n", articleId)
	// InsertArticle(article *model.ArticleDetail) (articleId int64, err error)

}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article success, len:%d\n", len(articleList))

}

func TestGetArticleDetail(t *testing.T) {
	articleDetail, err := GetArticleDetail(7)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article success, article:%#v\n", articleDetail)

}

func TestGetRelativeArticle(t *testing.T) {
	articleList, err := GetRelativeArticle(7)
	if err != nil {
		t.Errorf("get relative article failed, err:%v\n",err)
		return
	}

	for _, v := range articleList {
		t.Logf("id:%d title:%s\n",v.ArticleId,v.Title)
	}
}


func TestGetPrevArticleById(t *testing.T) {
	articleInfo, err := GetPrevArticleById(6)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n",err)
		return
	}
	t.Logf("article info:%#v\n",articleInfo)
}


func TestGetNextArticleById(t *testing.T) {
	articleInfo, err := GetNextArticleById(6)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n",err)
		return
	}
	t.Logf("article info:%#v\n",articleInfo)
}
