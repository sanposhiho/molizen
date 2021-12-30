package context

import (
	"sync"

	"github.com/sanposhiho/molizen/system"

	"github.com/sanposhiho/molizen/actor"
)

type Context struct {
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

func (c *Context) NewChildContext(
	actor actor.Actor,
	locker parentActorLocker,
	unlocker parentActorUnlocker,
) Context {
	c.system.RegisterActor(actor, c.parent)

	return Context{
		parent: &parent{
			locker:       locker,
			unlocker:     unlocker,
			isLockedByUs: true,
		},
	}
}

func (c *Context) UnlockParent() {
	if !c.hasParent() {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.parent.isLockedByUs {
		c.parent.unlocker()
		c.parent.isLockedByUs = false
		return
	}
}

func (c *Context) LockParent() {
	if !c.hasParent() {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.parent.isLockedByUs {
		c.parent.locker()
		c.parent.isLockedByUs = true
		return
	}
}

func (c *Context) hasParent() bool {
	return c.parent != nil
}
