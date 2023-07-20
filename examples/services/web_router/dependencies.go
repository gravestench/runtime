package web_router

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager != nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(manager runtime.R) {
	for _, other := range manager.Services() {
		if cfg, ok := other.(config_file.Manager); ok {
			s.cfgManager = cfg
		}
	}
}
