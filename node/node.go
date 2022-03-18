package node

import (
	"github.com/sanposhiho/molizen/actorlet"
	"github.com/sanposhiho/molizen/context"
)

type Node struct {
	actorlet *actorlet.ActorLet
}

func NewNode() *Node {
	sys := actorlet.NewActorLet()
	return &Node{
		actorlet: sys,
	}
}

func (n *Node) NewContext() context.Context {
	return context.NewInitialContext(n.actorlet)
}
