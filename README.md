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
	- [Background](#background)
		- [race condition and `sync` package in Go](#race-condition-and-sync-package-in-go)
			- [about `-race` option in Go](#about--race-option-in-go)
	- [Alternatives for actor-model in Go](#alternatives-for-actor-model-in-go)
		- [Sending messages explicitly](#sending-messages-explicitly)
		- [un-typed message passing](#un-typed-message-passing)
		- [no actor reentrancy](#no-actor-reentrancy)


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

```
-source: Input Go source file.
-destination: Output file; defaults to stdout.
-package: Package of the generated code; defaults to the package of the input with a 'actor_' prefix.
-copyright_file: Copyright file used to add copyright header
```

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
This is a big difference from other actor libraries in Go.

This allows you to use the benefits of object-oriented programming while working with actors.

### Actor reentrancy

Actors in Swift have the feature called "Actor reentrancy" and Molizen follow it.
> Actor-isolated functions are reentrant.
> When an actor-isolated function suspends, reentrancy allows other work to execute on the actor before the original actor-isolated function resumes, which we refer to as interleaving.
> Reentrancy eliminates a source of deadlocks, where two actors depend on each other, can improve overall performance by not unnecessarily blocking work on actors, and offers opportunities for better scheduling of (e.g.) higher-priority tasks.
> [0306 Actors # Actor reentrancy| apple/swift-evolution](https://github.com/apple/swift-evolution/blob/23405a18e3ebbe69fcb37b0d316aa4ec5a7b6c46/proposals/0306-actors.md#actor-reentrancy)

If no actor reentrancy, you need to be careful with deadlocks and it's too difficult to do that.
For example, if you write program which two actors send messages to each other,
both actors will wait for each other to finish processing the message, and a deadlock may occur.

You can see more detailed reason why actor reentrancy is needed in [0306 Actors # Actor reentrancy| apple/swift-evolution](https://github.com/apple/swift-evolution/blob/23405a18e3ebbe69fcb37b0d316aa4ec5a7b6c46/proposals/0306-actors.md#actor-reentrancy)

## What is actor-model?

The actor model is a concept developed in the paper "A Universal Modular ACTOR Formalism for Artificial Intelligence" written in 1973.

This architecture mainly handles an object called "Actor".
Actors are active objects that perform their roles according to defined behaviors,
and all operations for actors are performed by message-passing.

Each actor has its own queue where incoming messages are stored and actors retrieve messages one by one from the queue and process them.
One of the major advantages of the actor model is that it prevents multiple actions from being performed on a single actor at the same time,
thus it can prevent data races, etc and protect internal data.

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

This section compares Molizen with alternatives for actor-model in Go; protoactor-go and ergo.

- [asynkron/protoactor-go](https://github.com/asynkron/protoactor-go)
- [ergo-services/ergo](https://github.com/ergo-services/ergo)

This is a library for implementing the classic actor model in Go.

Let's look at these differences between Molizen and protoactor-go/ergo.

- Sending messages explicitly in protoactor-go.
- un-typed message passing in protoactor-go.
- no actor reentrancy in protoactor-go.

#### Sending messages explicitly

The way of communication between actors is very different between Molizen and protoactor-go/ergo.

In protoactor-go or ergo, you send messages to actors with `context.Send`(protoactor-go) or `process.Send`(ergo).
This is simple and easy to understand for those who are familiar with the actor-model and are used to programming with actor.

In Molizen, users does not directly send messages.
Communication between actors is done through method calls.
You can benefit from actor-model in a similar way to programming with normal struct.
This is easy to understand for users who are familiar with object-oriented programming.

And, it also gives you all the benefits of object-oriented programming
-- use an abstraction by `interface`, libraries for mocking, or etc.

#### un-typed message passing

protoactor-go and ergo don't support typed message passing.
You need to convert message one by one in actors.

One big problem is that you can send wrong type messages to actors. Here is the example written with protoactor-go.

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

The same issue also exists in ergo.

#### no actor reentrancy

No actor reentrancy in protoactor-go and ergo.

You need to be careful with deadlocks as described in [Actor reentrancy](#actor-reentrancy),
otherwise, deadlocks will happen to you.
