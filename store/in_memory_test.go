package store_test

import (
	"github.com/fatihkahveci/simple-matchmaking/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemory(t *testing.T) {
	inMemoryStore := store.NewInMemoryStore()
	user := store.User{
		ID:    "1",
		Score: 96,
	}

	//Test add user
	inMemoryStore.Add(user)
	getUser := inMemoryStore.Get(user.ID)

	assert.Equal(t, user.ID, getUser.ID)

	//Test Remove User
	user2 := store.User{
		ID:    "2",
		Score: 95,
	}

	inMemoryStore.Add(user2)
	inMemoryStore.Remove(user2)
	getUser2 := inMemoryStore.Get(user2.ID)

	assert.Equal(t, store.User{}, getUser2)

	//Test GetAll
	allUsers := inMemoryStore.GetAll()
	expectedUsers := make(map[string]store.User)
	expectedUsers[user.ID] = getUser
	assert.Equal(t, expectedUsers, allUsers)

	assert.Equal(t, "InMemory", inMemoryStore.GetName())

}
