package main

import (
	"fmt"

	actor_user "github.com/sanposhiho/molizen/playground/scenarios/scenario0/actor"

	"github.com/sanposhiho/molizen/node"

	"github.com/sanposhiho/molizen/context"
)

func main() {
	node := node.NewNode()
	ctx := node.NewContext()
	actorFuture := actor_user.New(&User{})
	actor := actorFuture.Get(ctx)

	// request actor to set age 1.
	future := actor.SetAge(ctx, 1)
	// wait actor to process the request.
	future.Get(ctx)

	// request actor to get age.
	future2 := actor.GetAge(ctx)

	// The age should be the one we requested.
	fmt.Println("[using actor] Result: ", future2.Get(ctx).Ret0)
}

type User struct {
	name string
	age  int
}

func (u *User) SetAge(ctx context.Context, age int) {
	u.age = age
}

func (u *User) GetAge(ctx context.Context) int {
	return u.age
}
