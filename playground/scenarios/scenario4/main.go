package main

import (
	"fmt"
	"log"

	"github.com/sanposhiho/molizen/actor"
	"github.com/sanposhiho/molizen/context"

	actor_user "github.com/sanposhiho/molizen/playground/scenarios/scenario4/actor"

	"github.com/sanposhiho/molizen/node"
)

func main() {
	node := node.NewNode()
	ctx := node.NewContext()
	actorFuture := actor_user.New(ctx, &User{}, actor.Option{})
	actor, err := actorFuture.Get(ctx).Actor, actorFuture.Get(ctx).Error
	if err != nil {
		log.Fatalf("failed to create actor: %v", err)
	}

	repo := context.ExtractActorRepo[*actor_user.UserActor](ctx)
	// try actorrepo.Get
	a, err := repo.Get(actor.ActorName())
	if err != nil {
		log.Fatalf("failed to get actor from repo: %v", err)
	}
	fmt.Print(a.ActorName())
}

type User struct {
	name string
	age  int
}

func (u *User) Say(ctx context.Context, msg string) {
	fmt.Println(msg)
}
