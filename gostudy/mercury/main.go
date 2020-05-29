package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/controller/account"
	"github.com/gnufree/gostudy/mercury/controller/answer"
	"github.com/gnufree/gostudy/mercury/controller/category"
	"github.com/gnufree/gostudy/mercury/controller/comment"
	"github.com/gnufree/gostudy/mercury/controller/favorite"
	"github.com/gnufree/gostudy/mercury/controller/question"
	"github.com/gnufree/gostudy/mercury/dal/db"
	"github.com/gnufree/gostudy/mercury/filter"
	"github.com/gnufree/gostudy/mercury/id_gen"
	maccount "github.com/gnufree/gostudy/mercury/middleware/account"
	"github.com/gnufree/gostudy/mercury/util"
)

func initTemplate(router *gin.Engine) {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.Static("/css/", "./static/css/")
	router.Static("/js/", "./static/js/")
	router.Static("/fonts/", "./static/fonts/")
	router.Static("/img/", "./static/img/")
}

func initDb() (err error) {
	dns := "root:123456@tcp(localhost:3306)/mercury?parseTime=true"
	err = db.Init(dns)

	if err != nil {
		return
	}
	return
}

func initFilter() (err error) {
	err = filter.Init("./data/filter.txt")
	if err != nil {
		logger.Error("init filter failed, err:%v", err)
		return
	}
	logger.Debug("init filter success.")
	return
}

func initSession() (err error) {
	err = maccount.InitSession("redis", "localhost:6379")
	//err = maccount.InitSession("memory", "")
	return
}

func main() {

	// 1. 创建路由器
	router := gin.Default()
	config := make(map[string]string)
	config["log_level"] = "debug"
	logger.InitLogger("console", config)

	err := initFilter()
	if err != nil {
		panic(err)
	}
	err = initDb()
	if err != nil {
		panic(err)
	}

	err = id_gen.Init(1)
	if err != nil {
		panic(err)
	}

	err = initSession()
	if err != nil {
		panic(err)
	}

	err = util.InitKafka("localhost:9092")
	if err != nil {
		panic(err)
	}
	err = util.InitKafkaConsumer("localhost:9092", "mercury_topic",
		func(message *sarama.ConsumerMessage) {
			logger.Debug("eceive from kafka, msg:%#v", message)
		})
	if err != nil {
		panic(err)
		return
	}
	ginpprof.Wrapper(router)
	//initTemplate(router)

	// 2. 注册路由
	router.POST("/api/user/register", account.RegisterHandle)
	// 登录接口
	router.POST("/api/user/login", account.LoginHandle)

	router.GET("/api/category/list", category.GetCategoryListHandle)

	router.POST("/api/ask/submit", maccount.AuthMiddleware, question.QuestionSubmitHandle)

	router.GET("/api/question/list", category.GetQuestionListHandle)
	router.GET("/api/question/detail", question.QuestionDetailHandle)
	router.GET("/api/answer/list", answer.AnswerListlHandle)

	// 评论模块
	commentGroup := router.Group("/api/comment/", maccount.AuthMiddleware)
	//commentGroup := router.Group("/api/comment/")
	commentGroup.POST("/post_comment", comment.PostCommentHandle)
	commentGroup.POST("/post_reply", comment.PostReplyHandle)
	commentGroup.GET("/list", comment.CommentListHandle)

	// 收藏模块
	favoriteGroup := router.Group("/api/favorite/", maccount.AuthMiddleware)
	//favoriteGroup := router.Group("/api/favorite/")
	favoriteGroup.POST("/add_dir", favorite.AddDirHandle)
	favoriteGroup.POST("/add", favorite.AddFavoriteHandle)
	favoriteGroup.GET("/dir_list", favorite.DirListHandle)
	favoriteGroup.GET("/list", favorite.FavoriteListHandle)
	router.Run(":9090")
}
