package serviceB

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/pkg"
)

type hasA interface{ A() string }

func New(name string) *Service {
	return &Service{
		name: name,
	}
}

type Service struct {
	log  *zerolog.Logger
	name string

	dependency hasA // depends on service B
}

func (s *Service) B() string {
	return "this message came from ServiceB"
}

func (s *Service) Init(r runtime.Runtime) {
	s.log.Info().Msgf("calling A(): %s", s.dependency.A())
	return
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.log = logger
}

func (s *Service) DependenciesResolved() bool {
	return s.dependency != nil
}

func (s *Service) ResolveDependencies(rt pkg.IsRuntime) {
	// here, we iterate over all services from the runtime
	// and check if the service implements something we need.
	for _, service := range rt.Services() {
		if a, ok := service.(hasA); ok {
			s.dependency = a // If we find our hasA, we assign it!
		}
	}
}
