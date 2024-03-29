package common

import "time"

type Question struct {
	QuestionId    int64     `json:"question_id" db:"question_id"`
	Caption       string    `json:"caption" db:"caption"`
	Content       string    `json:"content" db:"content"`
	AuthorId      int64     `json:"author_id" db:"author_id"`
	CategoryId    int64     `json:"category_id" db:"category_id"`
	Status        int32     `json:"status" db:"status"`
	CreateTime    time.Time `json:"-" db:"create_time"`
	CreateTimeStr string    `json:"create_time"`
}
type ApiQuestion struct {
	Question
	AuthorName string `json:"author_name" db:"username"`
}

type ApiQuestionDetail struct {
	Question
	AuthorName   string `json:"author_name" db"author_name"`
	CategoryName string `json:"category_name" db:"category_name"`
}
