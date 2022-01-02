// Code generated by Molizen. DO NOT EDIT.

// Package actor_user is a generated Molizen package.
package actor_user

import (
	sync "sync"

	context "github.com/sanposhiho/molizen/context"
	future "github.com/sanposhiho/molizen/future"
)

// UserActor is a actor of User interface.
type UserActor struct {
	lock     sync.Mutex
	internal User
}

type User interface {
	Say(ctx context.Context, msg string)
}

func New(internal User) *UserActor {
	return &UserActor{
		internal: internal,
	}
}

// SayResult is the result type for Say.
type SayResult struct {
}

// Say actor base method.
func (a *UserActor) Say(ctx context.Context, msg string) *future.Future[SayResult] {
	ctx.UnlockSender()
	newctx := ctx.NewChildContext(a, a.lock.Lock, a.lock.Unlock)

	f := future.New[SayResult]()
	go func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.internal.Say(newctx, msg)

		ret := SayResult{}

		ctx.LockSender()

		f.Send(ret)
	}()

	return f
}