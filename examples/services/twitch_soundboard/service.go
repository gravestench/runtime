package twitch_soundboard

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/desktop_notification"
	"github.com/gravestench/runtime/examples/services/twitch_integration"
)

// this is an example service that implements all handlers for the
// twitch client we are using

type recipe interface {
	runtime.IsRuntimeService
	runtime.HasLogger
	runtime.HasDependencies
	config_file.HasDefaultConfig
	twitch_integration.OnPrivateMessage
}

var _ recipe = &Service{}

type Service struct {
	configManager config_file.Manager // dependency on config file manager
	notification  desktop_notification.SendsNotifications
	log           *zerolog.Logger
}

func (s *Service) Init(r runtime.R) {
	// nothing to do
}

func (s *Service) Name() string {
	return "Twitch Chat Soundboard"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.log = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log
}
