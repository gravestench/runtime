package main

import (
	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	rt.Add(&example{name: "foo"})

	rt.Run()
}
