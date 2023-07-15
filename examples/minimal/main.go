package main

import (
	runtime "runtime/pkg"
)

func main() {
	rt := runtime.New()

	for _, service := range []runtime.RuntimeServiceInterface{
		&example{name: "foo"},
	} {
		rt.Add(service)
	}

	rt.Run()
}
