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
		locker senderActorLocker,
		unlocker senderActorUnlocker,
	) *context
	UnlockSender()
	LockSender()
}

type context struct {
	mu     sync.Mutex
	let    *actorlet.ActorLet
	repo   actorrepo.ActorRepo
	sender *sender
}

type sender struct {
	actor        actor.Actor
	locker       senderActorLocker
	unlocker     senderActorUnlocker
	isLockedByUs bool
}

type senderActorLocker func()
type senderActorUnlocker func()

func NewInitialContext(let *actorlet.ActorLet, repo actorrepo.ActorRepo) *context {
	return &context{
		let:  let,
		repo: repo,
	}
}

func (c *context) NewChildContext(
	actor actor.Actor,
	locker senderActorLocker,
	unlocker senderActorUnlocker,
) *context {

	// TODO: register actor to repo
	// if err := c.repo.Create(actor); err != nil {
	// }

	return &context{
		let:  c.let,
		repo: c.repo,
		sender: &sender{
			locker:       locker,
			unlocker:     unlocker,
			isLockedByUs: true,
		},
	}
}

func (c *context) UnlockSender() {
	if !c.hasSender() {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.sender.isLockedByUs {
		c.sender.unlocker()
		c.sender.isLockedByUs = false
		return
	}
}

func (c *context) LockSender() {
	if !c.hasSender() {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.sender.isLockedByUs {
		c.sender.locker()
		c.sender.isLockedByUs = true
		return
	}
}

func (c *context) hasSender() bool {
	return c.sender != nil
}
