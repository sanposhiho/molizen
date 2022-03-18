package actorlet

import (
	"fmt"
	"time"

	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/actorrepo"
)

type actorLet[T1 actor.Actor, T2 actorrepo.ActorRepo[T1]] struct {
	repo T2
}

func NewActorLet[T1 actor.Actor, T2 actorrepo.ActorRepo[T1]](repo T2) *actorLet[T1, T2] {
	return &actorLet[T1, T2]{
		repo: repo,
	}
}

func (l *actorLet[T1, T2]) Run() {
	for range time.Tick(3 * time.Millisecond) {
		fmt.Println("Tick!!")
	}
}
