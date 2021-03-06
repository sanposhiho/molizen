package group

import (
	"errors"
	"fmt"
	"github.com/sanposhiho/molizen/context"
	"sync"

	"github.com/sanposhiho/molizen/future"
)

type FutureGroup[T any] struct {
	futures map[string]*future.Future[T]
	mu      sync.Mutex
}

func NewFutureGroup[T any]() FutureGroup[T] {
	return FutureGroup[T]{
		futures: make(map[string]*future.Future[T]),
	}
}

func (f FutureGroup[T]) Register(fu *future.Future[T], key string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.futures[key] = fu
}

var ErrNotFound = errors.New("future is not found")

func (f FutureGroup[T]) Get(ctx context.Context, key string) (T, error) {
	fu, ok := f.futures[key]
	if !ok {
		var t T
		return t, fmt.Errorf("get a future, key: %v, err: %w", key, ErrNotFound)
	}

	return fu.Get(ctx), nil
}

func (f FutureGroup[T]) Wait(ctx context.Context) {
	wg := sync.WaitGroup{}
	for _, fu := range f.futures {
		fu := fu
		wg.Add(1)
		go func() {
			defer wg.Done()

			fu.Get(ctx)
		}()
	}

	wg.Wait()
	return
}
