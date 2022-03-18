package actorrepo

import "github.com/sanposhiho/molizen/actor"

type ActorRepo[T actor.Actor] interface {
	Get(actorName string) (T, error)
	Apply(actor T) (T, error)
	Delete(actorName string) error
}
