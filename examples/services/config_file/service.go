package config_file

import (
	"sync"

	ee "github.com/gravestench/eventemitter"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
)

const (
	defaultConfigDir  = "~/.config"
	defaultConfigFile = "config.json"
)

// Service is a config file manager that marshals to and from json files.
type Service struct {
	log                        *zerolog.Logger
	mux                        sync.Mutex
	events                     *ee.EventEmitter
	configs                    map[string]*Config
	servicesWithDefaultConfigs map[string]HasDefaultConfig
	dir                        string
}

// BindLogger satisfies the runtime.HasLogger interface
func (s *Service) BindLogger(l *zerolog.Logger) {
	s.log = l
}

// Logger satisfies the runtime.HasLogger interface
func (s *Service) Logger() *zerolog.Logger {
	return s.log
}

// Name satisfies the runtime.IsRuntimeService interface
func (s *Service) Name() string {
	return "Config File Manager"
}

// Init satisfies the runtime.IsRuntimeService interface
func (s *Service) Init(manager runtime.Runtime) {
	s.configs = make(map[string]*Config)
	s.servicesWithDefaultConfigs = make(map[string]HasDefaultConfig)

	go s.loopApplyDefaultConfigs(manager)
}
