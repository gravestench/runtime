package main

import (
	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	rt.Add(&sender{})
	rt.Add(&receiver{})

	rt.Run()
}
