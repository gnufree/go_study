package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, option ...string) (err error) {
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error){
	s.rwlock.RLock()
	defer s.rwlock.Unlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer  s.rwlock.Unlock()
	var id [16]byte
	id = uuid.NewV4()

	sessionId := id.String()
	session = NewMemorySession(sessionId)

	s.sessionMap[sessionId] = session
	return
}

