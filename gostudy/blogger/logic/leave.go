package logic

import (
	"fmt"
	"github.com/gnufree/gostudy/blogger/dal/db"
	"github.com/gnufree/gostudy/blogger/model"
)

func GetLeaveList() (LeaveList []*model.Leave, err error) {
	LeaveList, err = db.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave list failed, err:%v\n",err)
		return
	}
	return
}

func InsertLeave(username,email, content string) (err error) {
	var c model.Leave
	c.Username = username
	c.Email = email
	c.Content = content

	err = db.InsertLeave(&c)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v\n",err)
		return
	}
	return
}
