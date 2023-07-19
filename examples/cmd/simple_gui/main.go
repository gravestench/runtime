package main

import (
	"github.com/faiface/mainthread"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/raylib_renderer"
)

func main() {
	rt := runtime.New()

	rt.Add(&config_file.Service{RootDirectory: "~/.config/runtime/example/simple_gui"})
	rt.Add(&raylib_renderer.Service{})

	// create 100 layers, each will show a moving circle
	for i := 0; i < 100; i++ {
		rt.Add(&Layer{})
	}

	mainthread.Run(rt.Run)
}
