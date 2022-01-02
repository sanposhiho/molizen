package context

import (
	"sync"

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
	system *actorlet.ActorLet
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

func NewEmptyContext() *context {
	return &context{}
}

func (c *context) NewChildContext(
	actor actor.Actor,
	locker senderActorLocker,
	unlocker senderActorUnlocker,
) *context {
	c.system.RegisterActor(actor, c.sender)

	return &context{
		system: c.system,
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
