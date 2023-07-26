package runtime

import (
	"github.com/gravestench/runtime/pkg"
)

/*
	these are just some exports to:
	- prevent you from having to know to import from pkg
	- make the interfaces less wordy in your code
*/

type (
	Runtime = pkg.IsRuntime
	R       = Runtime // for even more brevity

	Service = pkg.IsRuntimeService
	S       = Service
)

// use these interfaces to build your runtime services
type (
	HasGracefulShutdown = pkg.HasGracefulShutdown
	HasLogger           = pkg.HasLogger
	HasDependencies     = pkg.HasDependencies
	UsesEventBus        = pkg.UsesEventBus
)

var New = pkg.New
var _ = New
