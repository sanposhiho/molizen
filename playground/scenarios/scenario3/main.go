package main

import (
	"fmt"
	"strconv"

	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/context"

	actor_user "github.com/sanposhiho/molizen/playground/scenarios/scenario3/actor"

	"github.com/sanposhiho/molizen/future/group"
	"github.com/sanposhiho/molizen/node"
)

func main() {
	node := node.NewNode()
	ctx := node.NewContext()
	actorFuture := actor_user.New(ctx, &User{}, actor.Option{})
	actor := actorFuture.Get(ctx).Actor

	g := group.NewFutureGroup[actor_user.SayResult]()
	for i := 0; i < 100; i++ {
		future := actor.Say(ctx, strconv.Itoa(i))
		g.Register(future, strconv.Itoa(i))
	}

	g.Wait(ctx)
}

type User struct {
	name string
	age  int
}

func (u *User) Say(ctx context.Context, msg string) {
	fmt.Println(msg)
}
