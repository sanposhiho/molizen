package actorlet

import (
	"fmt"

	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/actorrepo"
	"github.com/sanposhiho/molizen/event"
)

type ActorLet interface {
	Run()
	Stop()
}

type actorLet[T1 actor.Actor, T2 actorrepo.ActorRepo[T1]] struct {
	repo   T2
	stopCh chan struct{}
}

func NewActorLet[T1 actor.Actor, T2 actorrepo.ActorRepo[T1]](repo T2) *actorLet[T1, T2] {
	return &actorLet[T1, T2]{
		repo: repo,
	}
}

func (l *actorLet[T1, T2]) Run() {
	ch := l.repo.Watch().ResultChan()
	for {
		select {
		case v := <-ch:
			if v.Type == event.Added {
				actor := v.Actor
				fmt.Print(actor.ActorName())
			}
		case <-l.stopCh:
			return
		}
	}
}

func (l *actorLet[T1, T2]) Stop() {
	l.stopCh <- struct{}{}
}
