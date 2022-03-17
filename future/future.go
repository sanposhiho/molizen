package future

import (
	"sync"

	"github.com/sanposhiho/molizen/context"
)

type Future[T any] struct {
	ch           chan T
	result       *T
	senderLocker senderLocker
}

type senderLocker struct {
	mu           sync.Mutex
	isLockedByUs bool
}

func New[T any]() *Future[T] {
	return &Future[T]{
		ch: make(chan T, 1),
		senderLocker: senderLocker{
			isLockedByUs: true,
		},
	}
}

func (f *Future[T]) Send(val T) {
	f.ch <- val
}

func (f *Future[T]) Get(ctx context.Context) T {
	if f.result == nil {
		result := f.get(ctx)
		f.result = &result
	}
	return *f.result
}

func (f *Future[T]) get(ctx context.Context) T {
	for {
		select {
		case result := <-f.ch:
			f.lockSender(ctx)
			return result
		default:
			f.unlockSender(ctx)
		}
	}
}

func (f *Future[T]) unlockSender(ctx context.Context) {
	if !ctx.HasSender() {
		return
	}

	f.senderLocker.mu.Lock()
	defer f.senderLocker.mu.Unlock()

	if f.senderLocker.isLockedByUs {
		ctx.SenderUnlocker()()
		f.senderLocker.isLockedByUs = false
		return
	}
}

func (f *Future[T]) lockSender(ctx context.Context) {
	if !ctx.HasSender() {
		return
	}

	f.senderLocker.mu.Lock()
	defer f.senderLocker.mu.Unlock()

	if !f.senderLocker.isLockedByUs {
		ctx.SenderLocker()()
		f.senderLocker.isLockedByUs = true
		return
	}
}
