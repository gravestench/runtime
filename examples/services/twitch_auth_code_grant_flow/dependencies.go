package twitch_auth_code_grant_flow

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

func (s *Service) DependenciesResolved() bool {
	return s.cfg != nil
}

func (s *Service) ResolveDependencies(rt runtime.R) {
	for _, service := range rt.Services() {
		if candidate, ok := service.(config_file.Dependency); ok {
			s.cfg = candidate
		}
	}
}
