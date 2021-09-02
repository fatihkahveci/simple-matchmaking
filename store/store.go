package store

type Store interface {
	GetName() string
	Add(user User)
	Remove(user User)
	Get(id string) User
	GetAll() map[string]User
}
