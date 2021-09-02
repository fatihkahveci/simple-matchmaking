<p align="center">
    Simple rule based matchmaking for your online game with support of Redcon(RESP) protocol.
</p>


## 1- Simple Match Rule

Easiest usage of system is "Direct Match Rule" in this rule players match each other without any rule.

```bash
package main

import (
	"github.com/fatihkahveci/simple-matchmaking/matchmaking"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"time"
)

func main()  {
	inMemory := store.NewInMemoryStore()
	dur, _ := time.ParseDuration("10s")

	r := matchmaking.NewDirectMatchRule()


	respServer := server.NewRespServer(inMemory, ":1234")

	matcher := matchmaking.NewMatchmaking("test",respServer,inMemory, r, dur)

	matcher.Start()
}
```

## 2- Score Match Rule

In this example users match with given score rule which means match happens only user score 10 between 15.

```bash
package main

import (
	"github.com/fatihkahveci/simple-matchmaking"
	"github.com/fatihkahveci/simple-matchmaking/rules"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"time"
)

func main()  {
	inMemory := store.NewInMemoryStore()
	dur, _ :=time.ParseDuration("10s")

	r := rules.NewScoreMatchRule(10,15)


	respServer := server.NewRespServer(inMemory, ":1234")

	matcher := simpe_mm.NewMatchmaking("score",respServer,inMemory, r, dur)

	matcher.Start()
}
```

## 3- Custom Match Rule

In this example we create our own rule. We just need to follow "MatchRule" interface.

```
package main

import (
	"github.com/fatihkahveci/simple-matchmaking"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"time"
)


type CustomFieldMatchRule struct {
	Field string
	MinThreshold int
	MaxThreshold int
}

func NewCustomFieldMatchRule(field string,minThreshold, maxThreshold int) CustomFieldMatchRule {
	return CustomFieldMatchRule{
		Field: field,
		MinThreshold: minThreshold,
		MaxThreshold: maxThreshold,
	}
}

func (r CustomFieldMatchRule) Match(user1, user2 store.User) bool {
	user1Level := user1.Fields[r.Field].(int)
	user2Level := user2.Fields[r.Field].(int)


	minLevel := user1Level - r.MinThreshold
	maxLevel := user1Level + r.MaxThreshold

	if user2Level >= minLevel && user2Level <= maxLevel {
		return true
	}

	return false
}

func (r CustomFieldMatchRule) GetName() string {
	return "CustomField"
}

func main()  {
	inMemory := store.NewInMemoryStore()
	dur, _ :=time.ParseDuration("10s")

	r := NewCustomFieldMatchRule("level",10,20)


	respServer := server.NewRespServer(inMemory, ":1234")

	matcher := simpe_mm.NewMatchmaking("custom",respServer,inMemory, r, dur)

	matcher.Start()
}
```

## Usage

You can add new user to pool using any redis client. Thanks for Redcon ❤️

```
redis-cli -p 1234 add '{
  "id": "55",
  "score": 5
}'
```

Also you need to subscribe your matchmaking channel.

```
redis-cli -p 1234
127.0.0.1:1234> subscribe simple
Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "simple"
3) (integer) 1
1) "message"
2) "simple"
3) "{\"user_1\":{\"id\":\"12\",\"score\":5,\"join_time\":\"2021-08-31T22:57:27.109609+03:00\",\"fields\":null},\"user_2\":{\"id\":\"55\",\"score\":5,\"join_time\":\"2021-08-31T22:57:27.109614+03:00\",\"fields\":null},\"match_rule_name\":\"Direct\",\"time\":\"2021-08-31T22:57:27.116724+03:00\",\"action_type\":\"match\"}"

```

## TODOS

- More Server Support (WebSocket etc)
- More Store Support (Redis, Mongo etc)
- Multiple User Support (Team matchmaking)
