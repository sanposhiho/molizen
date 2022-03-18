package context

import (
	"sync"

	"github.com/sanposhiho/molizen/actorrepo"

	"github.com/sanposhiho/molizen/actorlet"

	"github.com/sanposhiho/molizen/actor"
)

type Context interface {
	NewChildContext(
		actor actor.Actor,
		locker func(),
		unlocker func(),
	) *context
	SenderLocker() func()
	SenderUnlocker() func()
	HasSender() bool
	setRepo(key, repo interface{})
	getRepo(key interface{}) interface{}
}

type context struct {
	mu     sync.Mutex
	let    *actorlet.ActorLet
	repos  map[any]any
	sender *sender
}

type sender struct {
	actor    actor.Actor
	locker   func()
	unlocker func()
}

func NewInitialContext(let *actorlet.ActorLet) *context {
	return &context{
		let:    let,
		sender: &sender{},
		repos:  map[any]any{},
	}
}

func (c *context) NewChildContext(
	actor actor.Actor,
	locker func(),
	unlocker func(),
) *context {

	return &context{
		let:   c.let,
		repos: c.repos,
		sender: &sender{
			locker:   locker,
			unlocker: unlocker,
		},
	}
}

func (c *context) SenderLocker() func() {
	return c.sender.locker
}

func (c *context) SenderUnlocker() func() {
	return c.sender.unlocker
}

func (c *context) HasSender() bool {
	return c.SenderUnlocker() != nil && c.SenderLocker() != nil
}

func (c *context) setRepo(key interface{}, repo interface{}) {
	_, ok := c.repos[key]
	if ok {
		// already exist
		return
	}

	c.repos[key] = repo
}

func (c *context) getRepo(key interface{}) interface{} {
	re, _ := c.repos[key]
	return re
}

func RegisterActorRepo[T actor.Actor, C Context](ctx C, repo actorrepo.ActorRepo[T]) {
	var key T
	ctx.setRepo(key, repo)
}

func ExtractActorRepo[T actor.Actor, C Context](ctx C) actorrepo.ActorRepo[T] {
	var key T
	repo := ctx.getRepo(key)
	typed, ok := repo.(actorrepo.ActorRepo[T])
	if !ok {
		return nil
	}

	return typed
}
