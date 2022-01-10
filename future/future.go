package future

import (
	"sync"
)

type Future[T any] struct {
	ch chan T
	result *T
	senderLocker *senderLocker
}

type senderLocker struct {
	mu     sync.Mutex
	locker       func()
	unlocker     func()
	isLockedByUs bool
}

func New[T any](
	locker func(),
	unlocker func(),
) *Future[T] {
	var sl *senderLocker
	if locker != nil && unlocker != nil {
		sl = &senderLocker{
			locker:       locker,
			unlocker:     unlocker,
			isLockedByUs: true,
		}
	}
	return &Future[T]{
		ch: make(chan T, 1),
		senderLocker: sl,
	}
}

func (f *Future[T]) Send(val T) {
	f.ch <- val
}

func (f *Future[T]) Get() T {
	if f.result == nil {
		result := f.get()
		f.result = &result
	}
	return *f.result
}

func (f *Future[T]) get() T {
	for  {
		select {
		case result := <-f.ch:
			f.lockSender()
			return result
		default:
			f.unlockSender()
		}
	}
}

func (f *Future[T]) unlockSender() {
	if !f.hasSender() {
		return
	}

	f.senderLocker.mu.Lock()
	defer f.senderLocker.mu.Unlock()

	if f.senderLocker.isLockedByUs {
		f.senderLocker.unlocker()
		f.senderLocker.isLockedByUs = false
		return
	}
}

func (f *Future[T]) lockSender() {
	if !f.hasSender() {
		return
	}

	f.senderLocker.mu.Lock()
	defer f.senderLocker.mu.Unlock()

	if !f.senderLocker.isLockedByUs {
		f.senderLocker.locker()
		f.senderLocker.isLockedByUs = true
		return
	}
}

func (f *Future[T]) hasSender() bool {
	return f.senderLocker != nil
}
