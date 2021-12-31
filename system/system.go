package system

import "github.com/sanposhiho/molizen/actor"

type ActorSystem struct {
	// TODO: The actor's structure will be recorded here as tree.
}

func NewActorSystem() *ActorSystem {
	return &ActorSystem{}
}

func (s *ActorSystem) RegisterActor(actor actor.Actor, sender actor.Actor) {
}
