package raylib_renderer

import (
	"os"
	"time"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/runtime"
)

func (s *Service) initRenderer() {
	cfg, err := s.Config()
	if err != nil {
		s.log.Error().Msgf("getting config: %v", err)
	}

	windowSettings := cfg.Group(keyGroupWindowSettings)
	screenWidth := windowSettings.GetInt(keyScreenWidth)
	screenHeight := windowSettings.GetInt(keyScreenHeight)

	// this requires runtime.Run() to be passed into mainthread.Run
	mainthread.Call(func() {
		rl.SetTargetFPS(60)
		rl.InitWindow(int32(screenWidth), int32(screenHeight), "runtime - raylib example")
	})
}

func (s *Service) gatherLayers(rt runtime.R) {
	for {
		for _, service := range rt.Services() {
			if _, alreadyBound := s.layers[service.Name()]; alreadyBound {
				continue
			}

			if candidate, ok := service.(RenderableLayer); ok {
				s.AddLayer(candidate)
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) renderServicesAsLayers(rt runtime.R) {
	s.mux.Lock()
	defer s.mux.Unlock()

	defer func() { _ = recover() /* don't worry about it */ }()

	mainthread.Call(func() {
		defer func() { _ = recover() /* don't worry about it */ }()

		for !rl.WindowShouldClose() {
			rl.BeginMode2D(rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 1))
			rl.BeginDrawing()
			rl.SetBlendMode(int32(rl.BlendAlpha))
			rl.ClearBackground(rl.Black)

			for _, name := range s.order {
				if layer := s.layers[name]; layer != nil {
					layer.OnRender()
				}
			}

			rl.EndDrawing()
		}

		rl.CloseWindow()
		s.log.Info().Msg("closing window")
		os.Exit(0)
	})
}
