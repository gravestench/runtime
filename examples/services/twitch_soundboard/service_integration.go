package twitch_soundboard

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/twitch_integration"
)

// Ensure that Service implements the required interfaces.
var (
	_ runtime.Service              = &Service{}
	_ runtime.HasLogger            = &Service{}
	_ runtime.HasDependencies      = &Service{}
	_ config_file.HasDefaultConfig = &Service{}
	_ IsTwitchSoundboard           = &Service{}
)

type IsTwitchSoundboard interface {
	twitch_integration.OnPrivateMessage
}
