package actorrepo

import "errors"

var (
	ErrUnexpectedActorType = errors.New("unexpected actor type")
	ErrNotFound            = errors.New("actor is not found")
)
