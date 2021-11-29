package actor

import "sync"

type Context struct {
	mu     sync.Mutex
	parent parent
}

type parent struct {
	locker         parentActorLocker
	unlocker       parentActorUnlocker
	isUnlockedOnce bool
}

type parentActorLocker func()
type parentActorUnlocker func()

func NewContext(
	locker parentActorLocker,
	unlocker parentActorUnlocker,
) Context {
	return Context{
		parent: parent{
			locker:         locker,
			unlocker:       unlocker,
			isUnlockedOnce: false,
		},
	}
}

func (c *Context) UnlockParent() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.parent.isUnlockedOnce {
		// parent actor is already unlocked once.
		return
	}
	c.parent.unlocker()
	return
}
