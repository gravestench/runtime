package main

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

func main() {
	rt := runtime.New()
	cfgManager := &config_file.Service{}

	rt.Add(cfgManager)

	// This service has a dependency on the config manager
	rt.Add(&serviceThatUsesConfigManager{})

	rt.Run()
}
