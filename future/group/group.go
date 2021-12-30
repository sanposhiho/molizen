package group

import (
	"sync"

	"github.com/sanposhiho/molizen/future"
)

type FutureGroup[T any] struct {
	futures map[string]future.Future[T]
	mu      sync.Mutex
}

func NewFutureGroup[T any]() FutureGroup[T] {
	return FutureGroup[T]{
		futures: make(map[string]future.Future[T]),
	}
}

func (f FutureGroup[T]) Register(fu future.Future[T], key string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.futures[key] = fu
}

func (f FutureGroup[T]) Get(key string) T {
	return f.futures[key].Get()
}

func (f FutureGroup[T]) Wait() {
	wg := sync.WaitGroup{}
	for _, fu := range f.futures {
		fu := fu
		wg.Add(1)
		go func() {
			defer wg.Done()

			fu.Get()
		}()
	}

	wg.Wait()
	return
}
