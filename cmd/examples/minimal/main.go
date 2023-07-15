package main

import (
	runtime "runtime/pkg"
)

func main() {
	rt := runtime.New()

	for _, service := range []runtime.RuntimeServiceInterface{
		&example{name: "foo"},
		&example{name: "bar"},
		&example{name: "baz"},
		&example{name: "bruh"},
		&example{name: "something"},
		&example{name: "nothing"},
		&example{name: "jamesies"},
		&example{name: "PackagedWhale"},
	} {
		rt.Add(service)
	}

	rt.Run()
}
