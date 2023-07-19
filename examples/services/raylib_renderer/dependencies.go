package raylib_renderer

import (
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/pkg"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	if cfg, err := s.Config(); cfg == nil || err != nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(isRuntime pkg.IsRuntime) {
	for _, service := range isRuntime.Services() {
		if candidate, ok := service.(config_file.Manager); ok {
			s.cfgManager = candidate
		}
	}
}
