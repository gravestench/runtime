package main

import (
	ee "github.com/gravestench/eventemitter"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/pkg"
)

type receiver struct {
	events *ee.EventEmitter
	logger *zerolog.Logger
}

func (r *receiver) BindLogger(logger *zerolog.Logger) {
	r.logger = logger
}

func (r *receiver) Logger() *zerolog.Logger {
	return r.logger
}

func (r *receiver) BindsEvents(emitter *ee.EventEmitter) {
	r.events = emitter
}

func (r *receiver) Init(rt pkg.IsRuntime) {
	r.events.On("test", func(args ...any) {
		r.logger.Info().Msgf("event 'test' recieved, args are: %+v", args)
	})
}

func (r *receiver) Name() string {
	return "receiver"
}
