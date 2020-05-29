package controller

import (
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/blogger/logic"
)

func IndexHandle(c *gin.Context) {

	artacleRecordList, err := logic.GetArticleRecordList(0, 15)
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	allCategoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n",err)
		return
	}
	var data map[string]interface{} = make(map[string]interface{},10)
	data["artacle_list"] = artacleRecordList
	data["category_list"] = allCategoryList
	c.HTML(http.StatusOK, "views/index.html", data)
}

func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr,10,64)

	if err != nil {
		fmt.Printf("strconv string to int64 failed, err:%v\n",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	artacleRecordList, err := logic.GetArticleRecordListById(int(categoryId),0,15)
	if err != nil {
		fmt.Printf("get article failed, err:%v\n",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	allCategoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n",err)
	}

	var data map[string]interface{} = make(map[string]interface{},10)
	data["artacle_list"] = artacleRecordList
	data["category_list"] = allCategoryList

	c.HTML(http.StatusOK,"views/index.html",data)

}

func NewArticle(c *gin.Context) {

	categoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}

func ArticleSubmit(c *gin.Context) {
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryIdStr := c.PostForm("category_id")
	title := c.PostForm("title")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		fmt.Printf("strconv string to int64 failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	err = logic.InsertArticle(content, author, title, categoryId)

	if err != nil {
		fmt.Printf("insert article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")

}

func ArticleDetail(c *gin.Context)  {

	articleIdStr := c.Query("article_id")
	articleId,err := strconv.ParseInt(articleIdStr,10,64)
	if err != nil {
		fmt.Printf("strconv string failed, err:%v\n",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}

	articleDetail, err := logic.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get articleDetail failed, err:%v\n",err)
		c.HTML(http.StatusInternalServerError, "views/500.html",nil)
		return
	}

	relativeArticle, err := logic.GetRelativeArticleList(articleId)
	if err != nil {
		fmt.Printf("get relative article failes, err:%v\n",err)
		return
	}

	prevArticle,nextArticle, err := logic.GetPrevAndNextArticleInfo(articleId)
	if err != nil {
		fmt.Printf("get prev or next article failed, err:%v\n",err)
	}

	categoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get all category list failed, err:%v\n",err)
		return
	}

	commentList, err := logic.GetCommentList(articleId)
	if err != nil {
		fmt.Printf("get comment list failed, err:%v\n",err)
	}
	err = logic.InsertViewCount(articleId)
	if err != nil {
		fmt.Printf("insert view count failed, err:%v\n",err)
	}

	fmt.Printf("relative article size:%d article_id:%d\n",len(relativeArticle),articleId)

	var m map[string]interface{} = make(map[string]interface{},10)
	m["detail"] = articleDetail
	m["relative_article"] = relativeArticle
	m["prev"] = prevArticle
	m["next"] = nextArticle
	m["category"] = categoryList
	m["article_id"] = articleId
	m["comment_list"] = commentList

	c.HTML(http.StatusOK,"views/detail.html",m)
}

func CommentSubmit(c *gin.Context)  {
	comment := c.PostForm("comment")
	username := c.PostForm("author")
	email := c.PostForm("email")
	articleIdStr := c.PostForm("article_id")
	artacleId, err := strconv.ParseInt(articleIdStr,10,64)
	if err != nil {
		fmt.Printf("strconv string to int64 failed, err:%v\n",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}

	err = logic.InsertComment(comment,username,email,artacleId)
	if err != nil {
		fmt.Printf("insert comment failed,err:%v\n",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	url := fmt.Sprintf("/article/detail/?article_id=%d",artacleId)
	c.Redirect(http.StatusMovedPermanently,url)


}

func LeaveSubmit(c *gin.Context)  {
	username := c.PostForm("author")
	content := c.PostForm("comment")
	email := c.PostForm("email")
	err := logic.InsertLeave(username,email,content)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v\n",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	url := fmt.Sprintf("/leave/new/")
	c.Redirect(http.StatusMovedPermanently,url)
}

func LeaveNew(c *gin.Context) {
	LeaveList,err := logic.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave list failed, err:%v\n",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	c.HTML(http.StatusOK,"views/gbook.html",LeaveList)
}

func AboutMe(c *gin.Context)  {
	c.HTML(http.StatusOK,"views/about.html",gin.H{
		"title":"posts",
	})
}