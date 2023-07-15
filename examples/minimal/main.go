package main

import (
	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	for _, service := range []runtime.Service{
		&example{name: "foo"},
	} {
		rt.Add(service)
	}

	rt.Run()
}
