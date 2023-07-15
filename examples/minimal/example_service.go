package main

import (
	"github.com/rs/zerolog"

	runtime "runtime/pkg"
)

type example struct {
	l    *zerolog.Logger
	name string
}

func (e *example) Init(r runtime.RuntimeInterface) {
	return
}

func (e *example) Name() string {
	return e.name
}

func (e *example) Logger() *zerolog.Logger {
	return e.l
}

func (e *example) UseLogger(logger *zerolog.Logger) {
	e.l = logger
}
