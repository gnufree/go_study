package account

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gnufree/gostudy/logger"
	"github.com/gnufree/gostudy/mercury/session"
)

func ProcessRequest(ctx *gin.Context) {

	var userSession session.Session
	defer func() {
		if userSession == nil {
			userSession = session.CreateSession()
		}
		ctx.Set(MercurySessionName, userSession)
	}()
	cookie, err := ctx.Request.Cookie(CookieSessionId)
	fmt.Printf("cookie:%v err:%v \n", cookie, err)
	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	sessionId := cookie.Value
	fmt.Printf("sessionId:%v\n", sessionId)
	if len(sessionId) == 0 {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	userSession, err = session.Get(sessionId)
	fmt.Printf("userSession:%#v\n", userSession)

	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))

		return
	}
	tmpUserId, err := userSession.Get(MercuryUserId)
	fmt.Printf("tmpUserId:%#v\n", tmpUserId)

	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	userId, ok := tmpUserId.(int64)
	fmt.Printf("userId:%#v\n", userId)

	if !ok || userId == 0 {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}

	ctx.Set(MercuryUserId, int64(userId))
	ctx.Set(MercuryUserLoginStatus, int64(1))
	return
}

func GetUserId(ctx *gin.Context) (userId int64, err error) {
	tempUerId, exists := ctx.Get(MercuryUserId)
	if !exists {
		err = errors.New("user id not exists")
		return
	}

	userId, ok := tempUerId.(int64)
	if !ok {
		err = errors.New("user id convert to int54 failed")
	}
	return

}

func IsLogin(ctx *gin.Context) (login bool) {
	tempLoginStatus, exists := ctx.Get(MercuryUserLoginStatus)
	if !exists {
		logger.Error("user not login temploginstatus:%v exists:%v", tempLoginStatus, exists)
		return
	}

	loginStatus, ok := tempLoginStatus.(int64)
	if !ok {
		return
	}

	if loginStatus == 0 {
		return
	}

	login = true
	return
}

func SetUserId(userId int64, ctx *gin.Context) {

	var userSession session.Session
	tempSession, exists := ctx.Get(MercurySessionName)
	if !exists {
		logger.Error("MercurySessionName not exists!")
		return
	}

	userSession, ok := tempSession.(session.Session)
	if !ok {
		logger.Error("tempSession not ok!")
		return
	}

	if userSession == nil {
		logger.Error("userSession eq nil")
		return
	}

	userSession.Set(MercuryUserId, userId)
}
func ProcessResponse(ctx *gin.Context) {

	var userSession session.Session
	tempSession, exists := ctx.Get(MercurySessionName)
	if !exists {
		return
	}

	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}

	if userSession.IsModify() == false {
		return
	}

	err := userSession.Save()
	if err != nil {
		return
	}

	sessionId := userSession.Id()
	cookie := &http.Cookie{
		Name:     CookieSessionId,
		Value:    sessionId,
		MaxAge:   CookieMaxAge,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(ctx.Writer, cookie)
	return

}
