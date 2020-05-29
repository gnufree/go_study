package logic

import (
	"fmt"

	"github.com/gnufree/gostudy/blogger/dal/db"
	"github.com/gnufree/gostudy/blogger/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	// 1 从数据库中，获取分类列表
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("1 get category list failed, err:%v\n", err)
		return
	}
	return
}
