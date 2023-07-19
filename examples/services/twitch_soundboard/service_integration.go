package twitch_soundboard

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/twitch_integration"
)

// this is a static check that my service satisfies the
// recipe below. This should prevent the code from compiling
// if this service should not implement these interfaces.
var _ recipe = &Service{}

type recipe interface {
	runtime.IsRuntimeService
	runtime.HasLogger
	runtime.HasDependencies
	config_file.HasDefaultConfig
	IsTwitchSoundboard
}

type IsTwitchSoundboard interface {
	twitch_integration.OnPrivateMessage
}
