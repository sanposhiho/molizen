package actor

import (
	"github.com/sanposhiho/molizen/actor"
	"sync"
	"context"

	"github.com/sanposhiho/molizen/future"

	"github.com/sanposhiho/molizen/playground/user"
)

type UserActor struct {
	internal user.User // 元の構造体
	mu       sync.Mutex
}

type SetNameResult struct {
	result1 string
}

func (u *UserActor) SetName(ctx actor.Context,name string) future.Future[string] {
	ctx.UnlockParent()
	ctx = actor.NewContext(u.mu.Lock, u.mu.Unlock)

	f := future.New[string]()
	go func() {
		u.mu.Lock()
		defer u.mu.Unlock()

		rto := u.internal.SetName(ctx, name)

		f.Send(rto)

		// ctxから呼び出しもとのActorの情報を取得して、
		// 何かしらを通して、その呼び出しもとのActorのメッセージの処理を中断する。
	}()

	return f
}

// generated

func (u *UserActor) SetAge(v int) future.Future[future.Result] {
	f := future.New[future.Result]()
	go func() {
		u.mu.Lock()
		defer u.mu.Unlock()

		u.internal.Age = v
		f.Send(future.Done)
	}()

	return f
}

func (u *UserActor) GetAge() future.Future[int] {
	f := future.New[int]()
	go func() {
		u.mu.Lock()
		defer u.mu.Unlock()

		v := u.internal.Age
		f.Send(v)
	}()

	return f
}
