package db

import (
	"database/sql"
	"fmt"

	"github.com/gnufree/gostudy/blogger/model"
	_ "github.com/go-sql-driver/mysql"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {

	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}

	sqlstr := `insert into
					article(content, summary, title, username,
						category_id,view_count,comment_count)
				values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary, article.Title,
		article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}

	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select 
					id, summary, title, view_count,
					create_time,comment_count,username,category_id
				from 
					article
				where
					status = 1
				order by create_time desc
				limit ?,?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

func GetArticleListByCategoryId(categoryId, pageNum,pageSize int) (artacleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d",pageNum,pageSize)
		return
	}

	sqlstr := `select 
					id,summary,title,view_count,
					create_time,comment_count,category_id
				from
						article
				where
					status = 1
					and
					category_id = ?
				order by create_time desc
				limit ?,?`
	err = DB.Select(&artacleList,sqlstr,categoryId,pageNum,pageSize)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter, article_id:%d\n", articleId)
		return
	}

	articleDetail = &model.ArticleDetail{}

	sqlstr := `select 
					id, summary, title, view_count,content,
					create_time,comment_count,username,category_id
				from 
					article
				where
					id = ?
				and
					status = 1
				`
	err = DB.Get(articleDetail, sqlstr, articleId)
	return
}

func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId,sqlstr,articleId)
	if err != nil {
		return
	}
	sqlstr = "select id,title from article where category_id=? and id != ? limit 10"
	err = DB.Select(&articleList,sqlstr,categoryId,articleId)
	return
}

func GetPrevArticleById(articleId int64) (info *model.RelativeArticle, err error) {

	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlstr := "select id, title from article where id < ? order by id desc limit 1"
	err = DB.Get(info, sqlstr, articleId)

	if err != nil {
		return
	}
	return

}

func GetNextArticleById(articleId int64) (info *model.RelativeArticle, err error) {

	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlstr := "select id, title from article where id > ? order by id asc limit 1"
	err = DB.Get(info, sqlstr, articleId)

	if err != nil {
		return
	}
	return

}

func IsArticleExist(articleId int64) (exists bool, err error) {
	var id int64
	sqlstr := "select id from article where id = ?"
	err = DB.Get(&id,sqlstr,articleId)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return

}

func InsertViewCount(articleId int64) (err error) {
	sqlstr := "update article set view_count = view_count + 1 where id = ?"
	_, err = DB.Exec(sqlstr,articleId)
	if err != nil {
		fmt.Printf("update article view_count failed, err:%v\n",err)
		return
	}
	return
}