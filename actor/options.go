package actor

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Option struct {
	ActorName string
}

func (o *Option) Complete() {
	if o.ActorName == "" {
		random, _ := uuid.NewRandom()
		// use random name.
		o.ActorName = random.String()
	}
}
