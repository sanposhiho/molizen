package user

import (
	"github.com/sanposhiho/molizen/context"
)

type User interface {
	SetAge(ctx context.Context, age int)
	IncrementAge(ctx context.Context)
	GetAge(ctx context.Context) int
	Say(ctx context.Context, msg string)
}
