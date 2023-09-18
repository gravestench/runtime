package pkg

import (
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestRuntime(t *testing.T) {
	rt := New()

	go func() {
		time.Sleep(time.Second * 3)
		rt.Shutdown().Wait()
	}()

	rt.Add(&exampleService{})

	rt.Run()
}

type exampleService struct {
	logger *zerolog.Logger
}

func (e *exampleService) BindLogger(logger *zerolog.Logger) {
	e.logger = logger
}

func (e *exampleService) Logger() *zerolog.Logger {
	return e.logger
}

func (e *exampleService) Init(rt IsRuntime) {

}

func (e *exampleService) Name() string {
	return "exmaple"
}

func (e *exampleService) OnShutdown() {
	time.Sleep(time.Second * 3)
	e.logger.Info().Msg("graceful shutdown completed")
}
