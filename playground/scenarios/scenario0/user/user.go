package user

import (
	"github.com/sanposhiho/molizen/context"
)

type User interface {
	SetAge(ctx context.Context, age int)
	GetAge(ctx context.Context) int
}
