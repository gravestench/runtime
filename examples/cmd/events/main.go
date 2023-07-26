package main

import (
	"github.com/gravestench/runtime"
)

const (
	eventFoo = "foo"
	eventBar = "bar"
)

func main() {
	rt := runtime.New()

	rt.Add(&sender{})
	rt.Add(&receiver{})

	rt.Run()
}
