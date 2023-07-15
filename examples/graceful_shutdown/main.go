package main

import (
	"time"

	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	for _, service := range []runtime.Service{
		&example{name: "foo"},
		&example{name: "bar"},
		&example{name: "baz"},
	} {
		rt.Add(service)
	}

	go func() {
		time.Sleep(time.Second * 3)
		rt.Shutdown()
	}()

	rt.Run()
}
