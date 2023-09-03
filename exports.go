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

	Service = interface {
		// Init initializes the service and establishes a connection to the
		// service IsRuntime.
		Init(rt Runtime)

		// Name returns the name of the service.
		Name() string
	}

	S = Service
)

// use these interfaces to build your runtime services
type (
	HasGracefulShutdown = pkg.HasGracefulShutdown
	HasLogger           = pkg.HasLogger
	HasDependencies     = pkg.HasDependencies

	EventHandlerServiceAdded                = pkg.EventHandlerServiceAdded
	EventHandlerServiceRemoved              = pkg.EventHandlerServiceRemoved
	EventHandlerServiceInitialized          = pkg.EventHandlerServiceInitialized
	EventHandlerServiceEventsBound          = pkg.EventHandlerServiceEventsBound
	EventHandlerServiceLoggerBound          = pkg.EventHandlerServiceLoggerBound
	EventHandlerRuntimeRunLoopInitiated     = pkg.EventHandlerRuntimeRunLoopInitiated
	EventHandlerRuntimeShutdownInitiated    = pkg.EventHandlerRuntimeShutdownInitiated
	EventHandlerDependencyResolutionStarted = pkg.EventHandlerDependencyResolutionStarted
	EventHandlerDependencyResolutionEnded   = pkg.EventHandlerDependencyResolutionEnded
)

var New = pkg.New
var _ = New
