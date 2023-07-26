package main

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

type recipe interface {
	runtime.R
	runtime.HasLogger
	runtime.HasDependencies
}

type serviceThatUsesConfigManager struct {
	configManager config_file.Manager // dependency on config file manager
	log           *zerolog.Logger
}

func (s *serviceThatUsesConfigManager) ResolveDependencies(rt runtime.R) {
	for _, service := range rt.Services() {
		if instance, ok := service.(config_file.Manager); ok {
			s.configManager = instance
		}
	}
}

func (s *serviceThatUsesConfigManager) DependenciesResolved() bool {
	return s.configManager != nil
}

func (s *serviceThatUsesConfigManager) Init(r runtime.R) {
	cfg, err := s.configManager.GetConfig("test.json")
	if err != nil {
		s.log.Fatal().Msgf("couldnt load example config file: %v", err)
	}

	group := cfg.Group("foo")

	group.Set("a", 1)
	group.Set("b", 2)
	group.Set("c", 3)

	s.configManager.SaveConfig("test.json")
}

func (s *serviceThatUsesConfigManager) Name() string {
	return "Config User"
}

func (s *serviceThatUsesConfigManager) BindLogger(logger *zerolog.Logger) {
	s.log = logger
}

func (s *serviceThatUsesConfigManager) Logger() *zerolog.Logger {
	return s.log
}
