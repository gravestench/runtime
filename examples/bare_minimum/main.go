package main

import (
	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New("my runtime")

	rt.Add(&example{name: "foo"})

	rt.Run()
}
