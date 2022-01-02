package actorlet

import "github.com/sanposhiho/molizen/actor"

type ActorLet struct{}

func NewActorLet() *ActorLet {
	return &ActorLet{}
}

func (s *ActorLet) RegisterActor(actor actor.Actor, sender actor.Actor) {
}
