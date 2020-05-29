package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/blogger/controller"
	"github.com/gnufree/gostudy/blogger/dal/db"
)

func main() {
	router := gin.Default()

	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)

	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexHandle)
	// 发布文章页面
	router.GET("/article/new/", controller.NewArticle)
	// 文章提交接口
	router.POST("/article/submit/", controller.ArticleSubmit)
	// 文章详情接口
	router.GET("/article/detail/", controller.ArticleDetail)
	// 分类下面的文章列表
	router.GET("/category/",controller.CategoryList)
	// 文章评论相关
	router.POST("/comment/submit/",controller.CommentSubmit)

	// 留言页面
	router.GET("/leave/new/",controller.LeaveNew)
	// 关于我页面
	router.GET("/about/me/",controller.AboutMe)
	// 留言提交接口
	router.POST("/leave/submit/",controller.LeaveSubmit)


	router.Run(":6060")

}
