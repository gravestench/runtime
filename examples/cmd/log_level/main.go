package main

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/pkg"
)

func main() {
	rt := runtime.New()

	rt.Add(&service{})

	rt.Run()
}

type service struct {
	logger *zerolog.Logger
}

func (s *service) Init(rt pkg.IsRuntime) {
	rt.SetLogLevel(zerolog.WarnLevel)
	s.logger.Trace().Msg("you should not see this")
	s.logger.Debug().Msg("you should not see this")
	s.logger.Info().Msg("you should not see this")

	rt.SetLogLevel(zerolog.InfoLevel)
	s.logger.Info().Msg("you should see this")
	s.logger.Debug().Msg("you should not see this")
	s.logger.Warn().Msg("you should see this")
	s.logger.Error().Msg("you should see this")

	rt.SetLogLevel(zerolog.TraceLevel)
	s.logger.Trace().Msg("you should see this")
	s.logger.Fatal().Msg("you should see this")
}

func (s *service) Name() string {
	return "Example"
}

func (s *service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *service) Logger() *zerolog.Logger {
	return s.logger
}
