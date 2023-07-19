package twitch_soundboard

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/desktop_notification"
)

// this is an example service that implements only the OnPrivateMessage handler

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
