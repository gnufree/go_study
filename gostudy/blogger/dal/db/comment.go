package db

import (
	"fmt"
	"github.com/gnufree/gostudy/blogger/model"
	_ "github.com/go-sql-driver/mysql"
)

func InsertComment(comment *model.Comment) (err error) {
	if comment == nil {
		err = fmt.Errorf("invalid parameter")
		return
	}
	tx, err := DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	sqlstr := `insert into comment(
					content, username, article_id)
				values (
					?,?,?
				)`

	_, err = tx.Exec(sqlstr,comment.Content,comment.Username,comment.ArticleId)
	if err != nil {
		return
	}
	sqlstr = "update article set comment_count = comment_count + 1 where id = ?"

	_, err = tx.Exec(sqlstr,comment.ArticleId)
	if err != nil {
		return
	}
	err = tx.Commit()
	return

}

func GetCommentList(artacleId int64, pageNum, pageSize int) (commentList []*model.Comment,err error ) {

	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid paramter, page_num:%d, page_size:%d",pageNum,pageSize)
		return
	}

	sqlstr := `select 
					id, content,username,create_time 
				from 
					comment
				where 
					article_id= ? 
				and
					status = 1
				order by create_time desc 
				limit ?,?`
	err = DB.Select(&commentList,sqlstr,artacleId,pageNum,pageSize)
	return
}
