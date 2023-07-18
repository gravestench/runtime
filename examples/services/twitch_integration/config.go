package twitch_integration

import (
	"github.com/gravestench/runtime/examples/services/config_file"
)

const (
	keyUsername = "Username"
	keyOauthKey = "Oauth Key"
)

func (s *Service) ConfigFilePath() string {
	return "twitch_integration.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	return s.cfgManager.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() (cfg config_file.Config) {
	g1 := cfg.Group("credentials")

	g1.Set(keyUsername, "your username")
	g1.Set(keyOauthKey, "your twitch oauth key")

	return
}
