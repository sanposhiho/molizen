# Molizen

**Molizen is not production-ready**

Molizen is a typed actor framework for Go.

This is a POWERFUL WEAPON to defeat "the difficulty of parallel programming" like race conditions and deadlocks.

- [Molizen](#molizen)
	- [Getting Started](#getting-started)
		- [0. install](#0-install)
		- [1. generate your Actor from interface](#1-generate-your-actor-from-interface)
			- [options](#options)
		- [2. use Actor](#2-use-actor)
	- [Design](#design)
		- [Actor reentrancy](#actor-reentrancy)
	- [What is actor-model?](#what-is-actor-model)
		- [Differences from CSP](#differences-from-csp)
		- [Benefits](#benefits)
	- [Background](#background)
		- [race condition and `sync` package in Go](#race-condition-and-sync-package-in-go)
			- [about `-race` option in Go](#about--race-option-in-go)
	- [Alternatives for actor-model in Go](#alternatives-for-actor-model-in-go)
		- [asynkron/protoactor-go](#asynkronprotoactor-go)
			- [Sending messages explicitly](#sending-messages-explicitly)
			- [un-typed message passing](#un-typed-message-passing)
		- [ergo-services/ergo](#ergo-servicesergo)


## Getting Started

Molizen requires Go 1.18+.

### 0. install

You need to install `molizen` command.

```shell
go install github.com/sanposhiho/molizen/cmd/molizen@latest
```

### 1. generate your Actor from interface

```shell
molizen [options]
```

This `molizen` command creates Actor struct from your interface.

For example, when you pass this `User` interface.

[/playground/scenarios/scenario0/user/user.go](/playground/scenarios/scenario0/user/user.go)

`molizen` command creates this `UserActor` struct.

[/playground/scenarios/scenario0/actor/user.go](/playground/scenarios/scenario0/actor/user.go)

#### options

**TBD**

### 2. use Actor

You can use generated `UserActor` like this.

```go
func main() {
	node := node.NewNode()
	ctx := node.NewContext()
	actorFuture := actor_user.New(ctx, &User{})
	actor := actorFuture.Get()

	// request actor to set age 1.
	future := actor.SetAge(ctx, 1)
	// wait actor to process the request.
	future.Get()

	// request actor to get age.
	future2 := actor.GetAge(ctx)

	// The age should be the one we requested.
	fmt.Println("[using actor] Result: ", future2.Get().Ret0)
}
```

You can run this code by using [/playground/scenarios/scenario0/main.go](/playground/scenarios/scenario0/main.go).
(See how to run scenarios in [/playground/scenarios/README.md](/playground/scenarios/README.md))

Some other examples of its use can be found [here](/playground/scenarios).

## Design

It is designed with reference to `actor` newly introduced in Swift5.5.

[0306 Actors | apple/swift-evolution](https://github.com/apple/swift-evolution/blob/23405a18e3ebbe69fcb37b0d316aa4ec5a7b6c46/proposals/0306-actors.md)

In Molizen, you can use actors like `struct` and communicate with other actors by its methods.
This is a big difference from other actor libraries.

This allows you to use the benefits of object-oriented programming while working with actors.

### Actor reentrancy

**TBD**

## What is actor-model?

**TBD**

### Differences from CSP

**TBD**

### Benefits

**TBD**

## Background

### race condition and `sync` package in Go

We usually avoid a race condition with `sync` package like this.

```go
mu := sync.Mutex{}
mu.Lock()
defer mu.Unlock()

// do something...
```

Even if you've been careful, you've probably encountered deadlocks and race conditions at one time or another.

Why? -- This is because any tools cannot find your race conditions and deadlocks completely.
You all have to avoid race conditions and deadlocks by yourself and no one doesn't support you.

As applications become more complex and multiple locks work in concert,
the level of difficulty becomes higher and higher.

Molizen is a powerful weapon and will help you to beat deadlocks and race conditions.
Let's change this hard game to an easier one with Molizen.

#### about `-race` option in Go

Yeah, maybe right now you're saying, "Go language has `-race` option to detect data race, doesn't it?"

`-race` option is awesome.
But, this is a tool that discovers race conditions at runtime, doesn't check all the code and find them.

> To start, run your tests using the race detector (go test -race).
> The race detector only finds races that happen at runtime,
> so it can't find races in code paths that are not executed.
> If your tests have incomplete coverage, you may find more races by running a binary built with -race under a realistic workload.
[Data Race Detector - The Go Programming Language](https://go.dev/doc/articles/race_detector#How_To_Use)

## Alternatives for actor-model in Go

This section compares Molizen with alternatives for actor-model in Go.

### asynkron/protoactor-go

[asynkron/protoactor-go](https://github.com/asynkron/protoactor-go)

> Proto Actor - Ultra fast distributed actors for Go, C# and Java/Kotlin

This is a library for implementing the classic actor model in Go.

Let's look at the differences between Molizen and protoactor-go.

#### Sending messages explicitly

The way of communication between actors is very different between Molizen and protoactor-go.

In protoactor-go, you send messages to actors with `context.Send`.

This is simple and easy to understand for those who are familiar with the Actor Model and are used to programming with it.

In Molizen, users does not directly send messages.

Communication between actors is done through method calls. Users can benefit from actor-model in a similar way to programming with normal struct.

This is easy to understand for users who are familiar with object-oriented programming.
It also gives you all the benefits of object-oriented programming
-- use an abstraction by `interface`, libraries for mocking, or etc.


#### un-typed message passing

protoactor-go doesn't support typed message passing.
You need to convert message one by one in actors.

One big problem is that you can send wrong type messages to actors.

```go
type Hello struct{ Who string }
type HelloV2 struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
	// convert received messages
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	// create actor
	sys := actor.NewActorSystem()
	context := actor.NewRootContext(sys, nil, opentracing.SenderMiddleware())
	props := actor.PropsFromProducer(func() actor.Actor { return &HelloActor{} })
	pid := context.Spawn(props)
	
	// correct message -- HelloActor can handle Hello message.
	context.Send(pid, Hello{Who: "Roger"}) 
	
	// incorrect message -- HelloActor cannot handle HelloV2 message.
	context.Send(pid, HelloV2{Who: "Roger"})

	time.Sleep(1 * time.Second)
}
```

This code passes compiling because any type of message can be sent.
You can't notice that You are sending the wrong message(`HelloV2`) at compile time.

### ergo-services/ergo

[ergo-services/ergo](https://github.com/ergo-services/ergo)

**TBD**