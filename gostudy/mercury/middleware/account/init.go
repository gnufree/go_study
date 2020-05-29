package account

import (
	"github.com/gnufree/gostudy/mercury/session"
)

func InitSession(provider string, addr string, options ...string) (err error)   {
	return session.Init(provider, addr, options...)
}
