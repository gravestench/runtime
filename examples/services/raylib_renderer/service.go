package raylib_renderer

import (
	"sync"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/pkg"
)

type Service struct {
	log        *zerolog.Logger
	mux        sync.Mutex
	cfgManager config_file.Manager // dependency
	layers     []RenderableLayer
}

func (s *Service) Init(rt pkg.IsRuntime) {
	defer func() { _ = recover() /* don't worry about it */ }()

	// raylib requires runtime.Run() to be passed into mainthread.Run in main.go
	go s.gatherLayers(rt)
	s.initRenderer()
	s.renderServicesAsLayers(rt)
}

func (s *Service) OnShutdown() {
	mainthread.Run(rl.CloseWindow)
}

func (s *Service) Name() string {
	return "Raylib Renderer"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.log = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log
}
