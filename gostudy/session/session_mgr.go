package session

type SessionMgr interface {
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session)
	Get(sessionId string) (session Session, err error)
}