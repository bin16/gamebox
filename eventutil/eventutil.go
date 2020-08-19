package eventutil

import (
	"log"
	"sync"
)

// namespace/session_id => channel

var namespaces = map[string]*Namespace{}
var rw = sync.RWMutex{}

// Namespace for Subscription
type Namespace struct {
	Name     string
	Sessions map[string]Session
}

// Session for Subscription
type Session struct {
	ID             string
	messageChannel chan interface{}
}

// OnMessage return message channel
func (s *Session) OnMessage() chan interface{} {
	return s.messageChannel
}

// Subscribe with Namespace and sessionID
func Subscribe(name string, sid string) *Session {
	log.Println("📢 Subscribe", name, "/", sid)
	rw.Lock()
	defer rw.Unlock()
	s1 := Session{
		ID:             sid,
		messageChannel: make(chan interface{}, 5),
	}
	if ns, ok := namespaces[name]; ok {
		if s, ok := ns.Sessions[sid]; ok {
			close(s.messageChannel)
			delete(ns.Sessions, sid)

			return &s
		}
		ns.Sessions[sid] = s1
		return &s1
	}

	namespaces[name] = &Namespace{
		Name: name,
		Sessions: map[string]Session{
			sid: s1,
		},
	}

	return &s1
}

// Post message into a NameSpace
func Post(name string, message interface{}) {
	rw.RLock()
	defer rw.RUnlock()
	log.Println("🚀 POST", name)
	if ns, ok := namespaces[name]; ok {
		for _, s := range ns.Sessions {
			ch := s.messageChannel
			go func() {
				ch <- message
			}()
		}
	}

}

// UnSubscribe clear things
func UnSubscribe(name, sid string) {
	rw.Lock()
	defer rw.Unlock()
	log.Println("🦖 UnSubscribe", name, "/", sid)
	if ns, ok := namespaces[name]; ok {
		if s, ok := ns.Sessions[sid]; ok {
			close(s.messageChannel)
			delete(ns.Sessions, sid)
			if len(ns.Sessions) == 0 {
				delete(namespaces, name)
			}
		}
	}
}
