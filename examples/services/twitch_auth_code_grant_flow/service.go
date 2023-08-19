package twitch_auth_code_grant_flow

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

type Service struct {
	log *zerolog.Logger
	cfg config_file.Dependency
	mux sync.Mutex

	stateString string

	token string
}

func (s *Service) Init(rt runtime.R) {

}

func (s *Service) Name() string {
	return "Twitch Integration (Auth code grant flow)"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.log = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log
}

func (s *Service) Token() (string, error) {
	if s.token == "" {
		return "", fmt.Errorf("not yet authorized")
	}

	return s.token, nil
}
