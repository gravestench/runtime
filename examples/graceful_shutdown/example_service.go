package main

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
)

type example struct {
	l    *zerolog.Logger
	name string
}

func (e *example) BindLogger(logger *zerolog.Logger) {
	e.l = logger
}

func (e *example) Logger() *zerolog.Logger {
	return e.l
}

func (e *example) Init(r runtime.Runtime) {
	return
}

func (e *example) Name() string {
	return e.name
}

func (e *example) OnShutdown() {
	e.l.Info().Msg("doing cleanup here...")
}
