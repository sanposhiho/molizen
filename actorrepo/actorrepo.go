package actorrepo

import "github.com/sanposhiho/molizen/actor"

type ActorRepo interface {
	Create(actor actor.Actor) error
}
