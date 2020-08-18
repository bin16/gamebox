package eventutil

import (
	"fmt"
	"sync"
)

// namespace/session_id => channel

var namespaces = map[string]*Namespace{}
var rw = sync.RWMutex{}

type Session struct {
	ID string
	Ch chan interface{}
}

type Namespace struct {
	Name     string
	Sessions []Session
}

func (ns *Namespace) getSession(sid string) Session {
	for _, s := range ns.Sessions {
		if s.ID == sid {
			return s
		}
	}

	return Session{}
}

func (ns *Namespace) containsSessionID(sid string) bool {
	for _, s := range ns.Sessions {
		if s.ID == sid {
			return true
		}
	}

	return false
}

func Subscribe(name string, sid string) chan interface{} {
	rw.Lock()
	defer rw.Unlock()
	if namespaces[name] == nil {
		namespaces[name] = &Namespace{Name: name}
	}
	ns := namespaces[name]
	if !ns.containsSessionID(sid) {
		ns.Sessions = append(ns.Sessions, Session{
			ID: sid,
			Ch: make(chan interface{}, 5),
		})
	}

	s := ns.getSession(sid)

	fmt.Println("Subscribe", name, "/", sid)

	return s.Ch
}

func Post(name string, message interface{}) {
	rw.RLock()
	defer rw.RUnlock()
	if namespaces[name] == nil {
		return
	}
	ns := namespaces[name]
	for _, s := range ns.Sessions {
		fmt.Println("POST to", name, "/", s.ID)
		s.Ch <- message
	}
}
