package session

import "fmt"

var (
	sessionMgr SessionMgr
)

func Init(provider string, addr string, options ...string) (err error) {

	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		err = fmt.Errorf("not support")
		return
	}
	err = sessionMgr.Init(addr,options...)
	return
}

func CreateSession() (session Session) {
	return sessionMgr.CreateSession()
}
func Get(sessionId string) (session Session, err error) {
	return sessionMgr.Get(sessionId)
}
