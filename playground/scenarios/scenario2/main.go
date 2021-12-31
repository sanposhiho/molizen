package main

import (
	"fmt"

	actor_user "github.com/sanposhiho/molizen/playground/scenarios/scenario2/actor"

	"github.com/sanposhiho/molizen/node"

	"github.com/sanposhiho/molizen/context"
)

func main() {
	node := node.NewNode()
	ctx := node.NewContext()
	actor := actor_user.New(&User{name: "taro"})
	actor2 := actor_user.New(&User{name: "hanako"})

	future := actor.SetSelf(ctx, actor)
	future.Get()
	future2 := actor2.SetSelf(ctx, actor2)
	future2.Get()

	future3 := actor.SendPing(ctx, actor2)

	future3.Get()
}

type User struct {
	name string
	age  int
	self *actor_user.UserActor
}

func (u *User) SetSelf(ctx context.Context, self *actor_user.UserActor) {
	u.self = self
}

func (u *User) Name(ctx context.Context) string {
	return u.name
}

func (u *User) SendPing(ctx context.Context, to *actor_user.UserActor) {
	future := to.Ping(ctx, u.self)

	future.Get()
}

func (u *User) Ping(ctx context.Context, from *actor_user.UserActor) {
	future := from.Name(ctx)

	name := future.Get().Ret0

	fmt.Printf("Hello %v\n", name)

	future2 := from.Pong(ctx)
	future2.Get()
	return
}

func (u *User) Pong(ctx context.Context) {
	fmt.Println("ponged")
}
