package user

import (
	"github.com/sanposhiho/molizen/actor"
)

type User interface {
	SetName(ctx actor.Context, name string) string
}
