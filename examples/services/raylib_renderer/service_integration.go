package raylib_renderer

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

// this is a static check that my service satisfies the
// recipe below. This should prevent the code from compiling
// if this service should not implement these interfaces.
var _ recipe = &Service{}

type recipe interface {
	runtime.Service
	runtime.HasLogger
	runtime.HasDependencies
	runtime.HasGracefulShutdown
	config_file.HasDefaultConfig
}

type RenderableLayer interface {
	runtime.S
	OnRender()
}
