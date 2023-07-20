package pkg

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// newLogger creates a factory function that generates a
// zerolog.Logger instance with a prefix and randomized colors based on the
// provided name. The luminosity delta between the colors is greater than 0.5.
func newLogger(service interface{ Name() string }) *zerolog.Logger {
	name := service.Name()

	writer := zerolog.ConsoleWriter{
		Out: os.Stdout,
		FormatMessage: func(input any) string {
			return fmt.Sprintf("[%s]: %s", name, input)
		},
	}

	logger := log.Output(writer).With().Logger()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &logger
}
