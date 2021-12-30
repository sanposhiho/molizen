package node

import (
	"github.com/sanposhiho/molizen/context"
	"github.com/sanposhiho/molizen/system"
)

type Node struct {
	system *system.ActorSystem
}

func NewNode() *Node {
	sys := system.NewActorSystem()
	return &Node{system: sys}
}

func (n *Node) NewContext() context.Context {
	return context.Context{}
}
