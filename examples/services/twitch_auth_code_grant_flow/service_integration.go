package twitch_auth_code_grant_flow

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/web_router"
)

type recipe interface {
	runtime.Service
	runtime.HasLogger
	runtime.HasDependencies
	config_file.HasDefaultConfig
	web_router.IsRouteInitializer
	TwitchAuthGrantFlow
}

var _ recipe = &Service{}

type TwitchAuthGrantFlow interface {
	Token() (string, error)
}
