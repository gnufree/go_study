package db

import (
	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/common"
)

func CreateQuestion(question *common.Question) (err error) {

	sqlstr := `insert into question(
								question_id,caption, content,author_id,category_id)
							values(?,?,?,?,?)`

	_, err = DB.Exec(sqlstr, question.QuestionId, question.Caption,
		question.Content, question.AuthorId, question.CategoryId)
	if err != nil {
		logger.Error("crate question failed, question:%#v, err:%v", question, err)
		return
	}

	return
}

func GetQuestionList(categoryId int64) (questionList []*common.Question, err error) {

	sqlstr := `select
					question_id,caption,content,author_id,category_id,create_time
				from
					question
				where category_id=?`
	err = DB.Select(&questionList, sqlstr, categoryId)
	if err != nil {
		logger.Error("get question list failed, err:%v", err)
		return
	}
	return

}

func GetQuestion(questionId int64) (question *common.Question, err error) {

	question = &common.Question{}

	sqlstr := `select
					question_id,caption,content,author_id,category_id,create_time
				from
					question
				where question_id=?`
	err = DB.Get(question, sqlstr, questionId)
	if err != nil {
		logger.Error("get question  failed, err:%v", err)
		return
	}
	return

}
