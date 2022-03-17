package context

import (
	"sync"

	"github.com/sanposhiho/molizen/actorrepo"

	"github.com/sanposhiho/molizen/actorlet"

	"github.com/sanposhiho/molizen/actor"
)

type Context interface {
	NewChildContext(
		actor actor.Actor,
		locker func(),
		unlocker func(),
	) *context
	SenderLocker() func()
	SenderUnlocker() func()
	HasSender() bool
}

type context struct {
	mu     sync.Mutex
	let    *actorlet.ActorLet
	repo   actorrepo.ActorRepo
	sender *sender
}

type sender struct {
	actor    actor.Actor
	locker   func()
	unlocker func()
}

func NewInitialContext(let *actorlet.ActorLet, repo actorrepo.ActorRepo) *context {
	return &context{
		let:    let,
		repo:   repo,
		sender: &sender{},
	}
}

func (c *context) NewChildContext(
	actor actor.Actor,
	locker func(),
	unlocker func(),
) *context {

	// TODO: register actor to repo
	// if err := c.repo.Create(actor); err != nil {
	// }

	return &context{
		let:  c.let,
		repo: c.repo,
		sender: &sender{
			locker:   locker,
			unlocker: unlocker,
		},
	}
}

func (c *context) SenderLocker() func() {
	return c.sender.locker
}

func (c *context) SenderUnlocker() func() {
	return c.sender.unlocker
}

func (c *context) HasSender() bool {
	return c.SenderUnlocker() != nil && c.SenderLocker() != nil
}
