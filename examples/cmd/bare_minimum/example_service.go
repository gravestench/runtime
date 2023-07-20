package main

import (
	"github.com/gravestench/runtime"
)

type example struct {
	name string
}

func (e *example) Init(r runtime.R) {
	return
}

func (e *example) Name() string {
	return e.name
}
