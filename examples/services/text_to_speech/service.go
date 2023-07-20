package text_to_speech

import (
	"fmt"
	"path/filepath"
	"time"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/pkg"
)

type Service struct {
	logger     *zerolog.Logger
	cfgManager config_file.Manager
	speech     htgotts.Speech
}

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	if cfg, _ := s.Config(); cfg == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(runtime pkg.IsRuntime) {
	for _, service := range runtime.Services() {
		if candidate, ok := service.(config_file.Manager); ok {
			s.cfgManager = candidate
		}
	}
}

func (s *Service) ConfigFilePath() string {
	return "text_to_speech.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfgManager.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() (cfg config_file.Config) {
	g := cfg.Group("Text to speech")

	cfgDir := s.cfgManager.ConfigDirectory()
	g.SetDefault("directory", filepath.Join(cfgDir, "audio_files"))

	return
}

func (s *Service) Init(rt pkg.IsRuntime) {
	var cfg *config_file.Config

	for { // wait until the config or default config is saved + loaded
		time.Sleep(time.Second)

		if cfg, _ = s.Config(); cfg != nil {
			break
		}
	}

	g := cfg.Group("Text to speech")

	s.speech = htgotts.Speech{Folder: g.GetString("directory"), Language: "en"}
}

func (s *Service) Name() string {
	return "Text-to-speech"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
