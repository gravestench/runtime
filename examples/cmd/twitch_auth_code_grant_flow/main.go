package main

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/twitch_auth_code_grant_flow"
	"github.com/gravestench/runtime/examples/services/web_router"
	"github.com/gravestench/runtime/examples/services/web_server"
)

func main() {
	rt := runtime.New()

	rt.Add(&config_file.Service{RootDirectory: "~/.config"})
	rt.Add(&web_server.Service{})
	rt.Add(&web_router.Service{})
	rt.Add(&twitch_auth_code_grant_flow.Service{})

	rt.Run()
}
