package main

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/text_to_speech"
)

func main() {
	rt := runtime.New()

	rt.Add(&config_file.Service{RootDirectory: "~/.config/runtime/example/text_to_speech"})
	rt.Add(&text_to_speech.Service{})
	rt.Add(&exampleServiceThatUsesTts{})

	rt.Run()
}
