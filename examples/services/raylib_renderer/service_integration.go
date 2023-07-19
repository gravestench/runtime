package raylib_renderer

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

// Ensure that Service implements the required interfaces.
var (
	_ runtime.Service              = &Service{}
	_ runtime.HasLogger            = &Service{}
	_ runtime.HasDependencies      = &Service{}
	_ runtime.HasGracefulShutdown  = &Service{}
	_ config_file.HasDefaultConfig = &Service{}
	_ LayerRenderer                = &Service{}
)

type LayerRenderer interface {
	AddLayer(RenderableLayer)
}

type RenderableLayer interface {
	OnRender()
}
