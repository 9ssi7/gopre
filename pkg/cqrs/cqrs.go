package cqrs

import "context"

// HandlerFunc is a function that can be used as a Handler
// It is used to convert a function to a Handler
type HandlerFunc[TParams any, TResponse any] func(context.Context, TParams) (TResponse, error)

// Handler is an interface that must be implemented by all command handlers
// It is used to execute a command
// The command handler must implement this interface
type Handler[TParams any, TResponse any] interface {
	Handle(context.Context, TParams) (TResponse, error)
}

// Empty is a type that can be used as a placeholder for empty cqrs responses
type Empty struct{}
