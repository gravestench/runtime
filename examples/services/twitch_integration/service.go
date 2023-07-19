package twitch_integration

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

type Service struct {
	logger     *zerolog.Logger
	cfgManager config_file.Manager

	twitchIrcClient *twitch.Client
}

func (s *Service) Init(rt runtime.R) {
	go s.setupClient()
	go s.loopBindHandlers(rt)
}

func (s *Service) Name() string {
	return "Twitch Integration"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
