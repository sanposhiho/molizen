package node

import (
	"github.com/sanposhiho/molizen/context"
)

type Node struct {
	actorlet any
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) NewContext() context.Context {
	return context.NewInitialContext(n.actorlet)
}
