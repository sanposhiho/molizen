package context

import (
	"sync"

	"github.com/sanposhiho/molizen/system"

	"github.com/sanposhiho/molizen/actor"
)

type Context interface {
	NewChildContext(
		actor actor.Actor,
		locker parentActorLocker,
		unlocker parentActorUnlocker,
	) *context
	UnlockParent()
	LockParent()
}

type context struct {
	mu     sync.Mutex
	system *system.ActorSystem
	parent *parent
}

type parent struct {
	actor        actor.Actor
	locker       parentActorLocker
	unlocker     parentActorUnlocker
	isLockedByUs bool
}

type parentActorLocker func()
type parentActorUnlocker func()

func NewEmptyContext() *context {
	return &context{}
}

func (c *context) NewChildContext(
	actor actor.Actor,
	locker parentActorLocker,
	unlocker parentActorUnlocker,
) *context {
	c.system.RegisterActor(actor, c.parent)

	return &context{
		system: c.system,
		parent: &parent{
			locker:       locker,
			unlocker:     unlocker,
			isLockedByUs: true,
		},
	}
}

func (c *context) UnlockParent() {
	if !c.hasParent() {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.parent.isLockedByUs {
		c.parent.unlocker()
		c.parent.isLockedByUs = false
		return
	}
}

func (c *context) LockParent() {
	if !c.hasParent() {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.parent.isLockedByUs {
		c.parent.locker()
		c.parent.isLockedByUs = true
		return
	}
}

func (c *context) hasParent() bool {
	return c.parent != nil
}
