package memory

import (
	"fmt"
	"sync"

	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/actorrepo"
)

type storage[T actor.Actor] struct {
	m sync.Map // map[actorname(string)]actor
}

func New[T actor.Actor]() actorrepo.ActorRepo[T] {
	return &storage[T]{}
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
	s.m.Store(actor.ActorName(), actor)
	return actor, nil
}

func (s *storage[T]) Delete(actorName string) error {
	s.m.Delete(actorName)
	return nil
}
