package actorrepo

import (
	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/event"
)

type ActorRepo[T actor.Actor] interface {
	Get(actorName string) (T, error)
	Apply(actor T) (T, error)
	Delete(actorName string) error
	Watch() EventWatcher[T]
}

type EventWatcher[T actor.Actor] interface {
	// Stops watching. Will close the channel returned by ResultChan(). Releases
	// any resources used by the watch.
	Stop()

	// Returns a chan which will receive all the events. If an error occurs
	// or Stop() is called, the implementation will close this channel and
	// release any resources used by the watch.
	ResultChan() <-chan event.Event[T]
}
