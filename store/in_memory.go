package store

import (
	"sync"
	"time"
)

type InMemory struct {
	Users map[string]User
	sync.Mutex
}

func NewInMemoryStore() Store {
	return &InMemory{
		Users: make(map[string]User),
	}
}

func (i *InMemory) Add(user User) {
	i.Lock()
	defer i.Unlock()
	user.JoinTime = time.Now()
	i.Users[user.ID] = user

}

func (i *InMemory) Remove(user User) {
	i.Lock()
	defer i.Unlock()
	delete(i.Users, user.ID)
}

func (i *InMemory) Get(id string) User {
	i.Lock()
	defer i.Unlock()
	return i.Users[id]
}

func (i *InMemory) GetAll() map[string]User {
	i.Lock()
	defer i.Unlock()
	return i.Users
}

func (i *InMemory) Len() int {
	i.Lock()
	defer i.Unlock()
	return len(i.Users)
}

func (i *InMemory) GetName() string {
	return "InMemory"
}
