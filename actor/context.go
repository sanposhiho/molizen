package actor

import "sync"

type Context struct {
	mu     sync.Mutex
	parent *parent
}

type parent struct {
	locker       parentActorLocker
	unlocker     parentActorUnlocker
	isLockedByUs bool
}

type parentActorLocker func()
type parentActorUnlocker func()

func NewEmptyContext() Context {
	return Context{}
}

func NewContext(
	locker parentActorLocker,
	unlocker parentActorUnlocker,
) Context {
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
