package pkg

import (
	"fmt"
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// newLogger is a factory function that generates a zerolog.Logger
func (r *Runtime) newLogger(service interface{ Name() string }, level zerolog.Level, dst io.Writer) *zerolog.Logger {
	name := service.Name()

	writer := zerolog.ConsoleWriter{
		Out: dst,
		FormatMessage: func(input any) string {
			return fmt.Sprintf("[%s]: %s", name, input)
		},
	}

	logger := log.Output(writer).With().Logger().Level(level)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &logger
}

func (r *Runtime) SetLogLevel(level zerolog.Level) {
	r.logger.Info().Msgf("setting log level to %s", level)

	r.logLevel = level

	// set the log-level for the runtime's logger
	r.logger = r.newLogger(r, r.logLevel, r.logOutput)

	// set the log level for each service that has a logger
	for _, service := range r.Services() {
		candidate, ok := service.(HasLogger)
		if !ok {
			continue
		}

		candidateLogger := r.newLogger(candidate, r.logLevel, r.logOutput)
		candidate.BindLogger(candidateLogger)
	}
}

func (r *Runtime) SetLogDestination(dst io.Writer) {
	r.logOutput = dst

	newLogger := r.newLogger(r, r.logLevel, r.logOutput)
	r.logger = newLogger

	// set the log level for each service that has a logger
	for _, service := range r.Services() {
		candidate, ok := service.(HasLogger)
		if !ok {
			continue
		}

		candidateLogger := r.newLogger(candidate, r.logLevel, r.logOutput)
		candidate.BindLogger(candidateLogger)
	}
}
