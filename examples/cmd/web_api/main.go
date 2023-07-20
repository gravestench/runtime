package main

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/cli_flags"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/web_router"
	"github.com/gravestench/runtime/examples/services/web_server"
)

func main() {
	rt := runtime.New()

	// will manage the config files for the other services
	rt.Add(&config_file.Service{})
	rt.Add(&cli_flags.Service{})
	rt.Add(&web_server.Service{})
	rt.Add(&web_router.Service{})

	rt.Add(&Service{}) // our example service that has routes

	rt.Run()
}
