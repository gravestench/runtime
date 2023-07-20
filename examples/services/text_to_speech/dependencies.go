package text_to_speech

import (
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/pkg"
)

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
