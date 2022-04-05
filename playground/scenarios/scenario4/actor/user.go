// Code generated by Molizen. DO NOT EDIT.

// Package actor_user is a generated Molizen package.
package actor_user

import (
	sync "sync"

	actor "github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/actorrepo/memory"
	context "github.com/sanposhiho/molizen/context"
	future "github.com/sanposhiho/molizen/future"
)

// UserActor is a actor of User interface.
type UserActor struct {
	name     string
	lock     sync.Mutex
	internal User
}

type User interface {
	Say(ctx context.Context, msg string)
}

// NewResult is the result type for New.
type NewResult struct {
	Actor UserActor
	// Error is an error that occurred during New.
	Error error
}

func New(ctx context.Context, internal User, opts actor.Option) *future.Future[NewResult] {
	opts.Complete()
	// TODO: make it selectable for users.
	repo := memory.New[*UserActor]()
	context.RegisterActorRepo(ctx, repo)
	f := future.New[NewResult]()
	go func() {
		actor := UserActor{
			internal: internal,
			name:     opts.ActorName,
		}
		_, err := repo.Apply(&actor)
		f.Send(NewResult{Actor: actor, Error: err})
	}()

	return f
}

// ActorName returns the actor's name.
func (a *UserActor) ActorName() string {
	return a.name
}

// SayResult is the result type for Say.
type SayResult struct {
}

// Say actor base method.
func (a *UserActor) Say(ctx context.Context, msg string) *future.Future[SayResult] {
	newctx := ctx.NewChildContext(a, a.lock.Lock, a.lock.Unlock)

	f := future.New[SayResult]()
	go func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.internal.Say(newctx, msg)

		ret := SayResult{}

		f.Send(ret)
	}()

	return f
}