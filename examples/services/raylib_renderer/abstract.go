package raylib_renderer

import (
	"github.com/gravestench/runtime"
)

type RenderableLayer interface {
	runtime.S
	OnRender()
}
