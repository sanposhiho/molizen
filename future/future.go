package future

type Future[T any] struct {
	ch chan T
}

func New[T any]() Future[T] {
	return Future[T]{
		ch: make(chan T),
	}
}

func (f Future[T]) Send(val T) {
	f.ch <- val
}

func (f Future[T]) Get() T {
	return <-f.ch
}
