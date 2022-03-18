package user

import (
	"github.com/sanposhiho/molizen/context"
)

type User interface {
	Say(ctx context.Context, msg string)
}
