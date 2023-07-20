package main

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/text_to_speech"
	"github.com/gravestench/runtime/examples/services/twitch_integration"
)

func main() {
	rt := runtime.New()

	rt.Add(&config_file.Service{RootDirectory: "~/.config/runtime/example/twitch_integrated_text_to_speech"})
	rt.Add(&twitch_integration.Service{})
	rt.Add(&text_to_speech.Service{})
	rt.Add(&glueService{})

	rt.Run()
}
