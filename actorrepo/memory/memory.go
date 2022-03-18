package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/actorrepo"
	"github.com/sanposhiho/molizen/event"
	"github.com/sanposhiho/molizen/event/manager"
)

type storage[T actor.Actor] struct {
	m        sync.Map // map[actorname(string)]actor
	eventMng manager.EventManager[T]
}

func New[T actor.Actor]() actorrepo.ActorRepo[T] {
	return &storage[T]{
		eventMng: manager.New[T](),
	}
}

func (s *storage[T]) Get(actorName string) (T, error) {
	a, ok := s.m.Load(actorName)
	if !ok {
		var zero T
		return zero, fmt.Errorf("load %v: %w", actorName, actorrepo.ErrNotFound)
	}

	typed, ok := a.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("expected %T type, but actually %v is returned from storage: %w", zero, a, actorrepo.ErrUnexpectedActorType)
	}

	return typed, nil
}
func (s *storage[T]) Apply(actor T) (T, error) {
	eve := event.Modified
	if _, err := s.Get(actor.ActorName()); err != nil {
		if !errors.Is(err, actorrepo.ErrNotFound) {
			var zero T
			return zero, fmt.Errorf("fetch actor to check the current status of actor: %w", err)
		}
		eve = event.Added

	}

	s.m.Store(actor.ActorName(), actor)
	s.eventMng.Publish(event.Event[T]{
		Type:  eve,
		Actor: actor,
	})

	return actor, nil
}

func (s *storage[T]) Delete(actorName string) error {
	s.m.Delete(actorName)
	return nil
}

func (s *storage[T]) Watch() actorrepo.EventWatcher[T] {
	return s.eventMng
}
