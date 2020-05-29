package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/common"
)

func CreatePostComment(comment *common.Comment) (err error) {
	tx, err := DB.Beginx()
	if err != nil {
		logger.Error("create post comment failed, comment:%#v, err:%v", comment, err)
		return
	}

	sqlstr := `insert
					into comment(
							comment_id,content,author_id,like_count,comment_count
					)
				values (
						?,?,?
				)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.Content, comment.AuthorId)
	if err != nil {
		logger.Error("insert into comment failed, comment:%#V, err", comment, err)
		tx.Rollback()
		return
	}

	sqlstr = `insert
					into comment_rel(
						comment_id,parent_id,level,
						question_id,reply_author_id
					)values (
						?,?,?,?,?
					)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.ParentId, 1,
		comment.QuestionId, comment.ReplyAuthorId)
	if err != nil {
		logger.Error("insert comment_rel failed, comment:%#v, err:%v", comment, err)
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		logger.Error("insert comment failed, comment:%#v, err:%v", comment, err)
		tx.Rollback()
		return
	}
	return
}

func CreateReplyComment(comment *common.Comment) (err error) {
	tx, err := DB.Beginx()
	if err != nil {
		logger.Error("create post comment failed, comment:%#v, err:%v", comment, err)
		return
	}
	// 1、根据ReplyCommentId查询对应的authorid
	var replyAuthorId int64
	sqlstr := `select author_id from comment where comment_id=?`
	err = tx.Get(&replyAuthorId, sqlstr, comment.ReplyCommentId)
	if err != nil {
		logger.Error("select author_id failed, err:%v, cid:%v", err, comment.ReplyCommentId)
		return
	}

	if replyAuthorId == 0 {
		err = fmt.Errorf("invalid reply author_id")
		return
	}
	comment.ReplyAuthorId = replyAuthorId

	sqlstr = `insert
					into comment(
							comment_id,content,author_id,like_count,comment_count
					)
				values (
						?,?,?
				)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.Content, comment.AuthorId)
	if err != nil {
		logger.Error("insert into comment failed, comment:%#V, err", comment, err)
		tx.Rollback()
		return
	}

	sqlstr = `insert
					into comment_rel(
						comment_id,parent_id,level,
						question_id,reply_author_id,reply_comment_id
					)values (
						?,?,?,?,?
					)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.ParentId, 2,
		comment.QuestionId, comment.ReplyAuthorId, comment.ReplyCommentId)
	if err != nil {
		logger.Error("insert comment_rel failed, comment:%#v, err:%v", comment, err)
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		logger.Error("insert comment failed, comment:%#v, err:%v", comment, err)
		tx.Rollback()
		return
	}
	return
}

func GetCommentList(answerId int64, offset, limit int64) (commentList []*common.Comment, count int64, err error) {

	//var commentList []int64
	sqlstr := `select comment_id from comment_rel where question_id=? and level=1 limit ?,?`
	err = DB.Select(&commentList, sqlstr, answerId, offset, limit)
	if err != nil {
		logger.Error("query comment list failed, answerid:%v, err:%v", answerId, err)
		return
	}

	if len(commentList) == 0 {
		return
	}

	sqlstr = `select 
					comment_id,content,author_id,like_count,comment_count,
					create_time
				from comment where comment_id in (?)`

	var interfaceSlice []interface{}
	for _, c := range commentList {
		interfaceSlice = append(interfaceSlice, c)
	}

	sqlstr, paramsList, err := sqlx.In(sqlstr, interfaceSlice)
	if err != nil {
		logger.Error("sqlx.In failed, answerid:%v, err:%v", answerId, err)
		return
	}
	err = DB.Select(&commentList, sqlstr, paramsList...)
	if err != nil {
		logger.Error("sql select failed, answerid:%v, err:%v", answerId, err)
		return
	}
	// 查询记录条数
	sqlstr = `select count(comment_id) from comment_rel where question_id=? and level=1`
	err = DB.Get(&count, sqlstr, answerId)
	if err != nil {
		logger.Error("query comment count failed, answerid:%v, err:%v", answerId, err)
		return
	}
	return
}
