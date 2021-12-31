package user

import (
	"github.com/sanposhiho/molizen/context"

	actor_user "github.com/sanposhiho/molizen/playground/scenarios/scenario2/actor"
)

type User interface {
	Name(ctx context.Context) string
	SendPing(ctx context.Context, to *actor_user.UserActor)
	Ping(ctx context.Context, from *actor_user.UserActor)
	Pong(ctx context.Context)
	SetSelf(ctx context.Context, self *actor_user.UserActor)
}
