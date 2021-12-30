package main

import (
	"fmt"

	"github.com/sanposhiho/molizen/actor"
	actor_user "github.com/sanposhiho/molizen/playground/actor"
)

func main() {
	ctx := actor.NewEmptyContext()
	actor := actor_user.New(&User{})
	future := actor.SetName(ctx, "sanposhiho")

	// get the result from future.
	result := future.Get()
	fmt.Println(result)
}

type User struct {
	name string
}

func (u *User) SetName(_ actor.Context, name string) string {
	u.name = name

	return name
}
