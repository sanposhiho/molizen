package event

import "github.com/sanposhiho/molizen/actor"

// Event represents a single event to a watched resource.
// +k8s:deepcopy-gen=true
type Event[T actor.Actor] struct {
	Type EventType

	Actor T
}

// EventType defines the possible types of events.
type EventType string

const (
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
	Error    EventType = "ERROR"
)
