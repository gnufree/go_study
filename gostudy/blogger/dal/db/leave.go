package db

import (
	"fmt"
	"github.com/gnufree/gostudy/blogger/model"
	_ "github.com/go-sql-driver/mysql"
)

func GetLeaveList() (leaveList []*model.Leave, err error) {
	sqlstr := "select id,username, email,content,create_time from `leave` order by id desc"

	err = DB.Select(&leaveList, sqlstr)
	if err != nil {
		fmt.Printf("exec sql:%s failed,err:%v\n",sqlstr,err)
		return
	}
	return
}

func InsertLeave(leave *model.Leave) (err error) {
	sqlstr := "insert into `leave`(username,email,content)values(?,?,?)"

	_, err = DB.Exec(sqlstr,leave.Username,leave.Email,leave.Content)
	if err != nil {
		fmt.Printf("exec sql:%s failed,err:%v\n",sqlstr,err)
		return
	}
	return

}