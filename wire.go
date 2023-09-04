//go:build wireinject

package main

import (
	"github.com/alewkinr/example-google-wire/internal"
	"github.com/google/wire"
)

func InitializeEvent() (internal.Event, error) {
	wire.Build(internal.NewEvent, internal.NewGreeter, internal.NewMessage)
	return internal.Event{}, nil
}
