package node

import (
	"github.com/sanposhiho/molizen/actorlet"
	"github.com/sanposhiho/molizen/context"
)

type Node struct {
	system *actorlet.ActorLet
}

func NewNode() *Node {
	sys := actorlet.NewActorLet()
	return &Node{system: sys}
}

func (n *Node) NewContext() context.Context {
	return context.NewEmptyContext()
}
