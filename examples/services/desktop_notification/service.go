package desktop_notification

import (
	"time"

	"github.com/gen2brain/beeep"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/pkg"
)

type Service struct {
	logger     *zerolog.Logger
	cfgManager config_file.Manager
}

func (s *Service) Init(rt pkg.IsRuntime) {

}

func (s *Service) Name() string {
	return "Desktop Notifications"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}

func (s *Service) DependenciesResolved() bool {
	return s.cfgManager != nil
}

func (s *Service) ResolveDependencies(r pkg.IsRuntime) {
	for _, candidate := range r.Services() {
		if service, ok := candidate.(config_file.Manager); ok {
			s.cfgManager = service
		}
	}
}

func (s *Service) ConfigFilePath() string {
	return "desktop_notifications.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	return s.cfgManager.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() (cfg config_file.Config) {
	g := cfg.Group("Notifications")

	g.SetDefault("timeout", time.Second*5)

	return cfg
}

func (s *Service) Notify(title, message, appIcon string) {
	err := beeep.Notify(title, message, appIcon)
	if err != nil {
		s.logger.Error().Msgf("sending notification: %v", err)
	}
}

func (s *Service) Alert(title, message, appIcon string) {
	err := beeep.Alert(title, message, appIcon)
	if err != nil {
		s.logger.Error().Msgf("sending notification: %v", err)
	}
}
