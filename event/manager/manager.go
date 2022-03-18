package manager

import (
	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/event"
)

type EventManager[T actor.Actor] interface {
	// Stops watching. Will close the channel returned by ResultChan(). Releases
	// any resources used by the watch.
	Stop()

	// Returns a chan which will receive all the events. If an error occurs
	// or Stop() is called, the implementation will close this channel and
	// release any resources used by the watch.
	ResultChan() <-chan event.Event[T]

	Publish(val event.Event[T])
}

type eventManager[T actor.Actor] struct {
	ch       chan event.Event[T]
	children []chan event.Event[T]
}

func New[T actor.Actor]() EventManager[T] {
	return &eventManager[T]{
		ch: make(chan event.Event[T]),
	}
}

func (w *eventManager[T]) Stop() {
	close(w.ch)
	for _, c := range w.children {
		close(c)
	}
}

func (w *eventManager[T]) ResultChan() <-chan event.Event[T] {
	newCh := make(chan event.Event[T])
	w.children = append(w.children, newCh)

	go func() {
		for {
			select {
			case e, ok := <-w.ch:
				if !ok {
					// w.ch is closed, so newCh will be closed soon.
					return
				}
				// send to channel
				newCh <- e
			}
		}
	}()

	return newCh
}

func (w *eventManager[T]) Publish(val event.Event[T]) {
	go func() {
		w.ch <- val
	}()
}
