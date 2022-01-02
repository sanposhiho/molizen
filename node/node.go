package node

import (
	"github.com/sanposhiho/molizen/actorlet"
	"github.com/sanposhiho/molizen/actorrepo"
	"github.com/sanposhiho/molizen/context"
)

type Node struct {
	actorlet  *actorlet.ActorLet
	actorrepo actorrepo.ActorRepo
}

func NewNode() *Node {
	sys := actorlet.NewActorLet()
	return &Node{
		actorlet: sys,
		// TODO: initialize actor repo.
		actorrepo: nil,
	}
}

func (n *Node) NewContext() context.Context {
	return context.NewInitialContext(n.actorlet, n.actorrepo)
}
