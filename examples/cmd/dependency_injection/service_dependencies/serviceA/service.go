package serviceA

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/pkg"
)

type hasB interface{ B() string }

func New(name string) *Service {
	return &Service{
		name: name,
	}
}

type Service struct {
	log  *zerolog.Logger
	name string

	dependency hasB // depends on service B
}

func (s *Service) A() string {
	return "this message came from ServiceA"
}

func (s *Service) Init(r runtime.Runtime) {
	s.log.Info().Msgf("calling B(): %s", s.dependency.B())
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
		if b, ok := service.(hasB); ok {
			s.dependency = b // If we find our hasB, we assign it!
			break
		}
	}
}
