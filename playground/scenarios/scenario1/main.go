package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/sanposhiho/molizen/future/group"
	"github.com/sanposhiho/molizen/node"

	"github.com/sanposhiho/molizen/context"

	actor_user "github.com/sanposhiho/molizen/playground/actor"
)

func main() {
	actormain()
	nonactormain()
}

func nonactormain() {
	node := node.NewNode()
	ctx := node.NewContext()
	user := User{}

	user.SetAge(ctx, 0)

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			user.IncrementAge(ctx)
		}()
	}

	wg.Wait()

	age := user.GetAge(ctx)
	fmt.Println("[using struct] Result: ", age)
}

func actormain() {
	node := node.NewNode()
	ctx := node.NewContext()
	actor := actor_user.New(&User{})
	future := actor.SetAge(ctx, 0)
	// wait to set age
	future.Get()

	g := group.NewFutureGroup[actor_user.IncrementAgeResult]()
	for i := 0; i < 100; i++ {
		future := actor.IncrementAge(ctx)
		g.Register(future, strconv.Itoa(i))
	}

	g.Wait()

	future2 := actor.GetAge(ctx)
	fmt.Println("[using struct] Result: ", future2.Get().Ret0)
}

type User struct {
	name string
	age  int
}

func (u *User) SetAge(ctx context.Context, age int) {
	u.age = age
}

// IncrementAge increment user's age.
// Note: not thread safe.
func (u *User) IncrementAge(ctx context.Context) {
	age := u.age
	u.age = age + 1
}

func (u *User) GetAge(ctx context.Context) int {
	return u.age
}
