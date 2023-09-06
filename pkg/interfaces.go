package pkg

import (
	"io"

	ee "github.com/gravestench/eventemitter"
	"github.com/rs/zerolog"
)

// IsRuntime is the abstract idea of the runtime, an interface.
//
// The IsRuntime interface defines the operations that can be performed on
// services, such as adding, removing, and retrieving services. It acts as a
// container for services and uses other interfaces like HasDependencies to
// work with them and do things automatically on their behalf.
type IsRuntime interface {
	// Add a single service to the IsRuntime.
	Add(IsRuntimeService)

	// Remove a specific service from the IsRuntime.
	Remove(IsRuntimeService)

	// Services returns a pointer to a slice of interfaces representing the
	// services currently managed by the service IsRuntime.
	Services() []IsRuntimeService

	SetLogLevel(level zerolog.Level)
	SetLogDestination(dst io.Writer)

	Events() *ee.EventEmitter

	Shutdown()
}

// IsRuntimeService represents a generic service within a runtime.
//
// The IsRuntimeService interface defines the contract that all services in the
// runtime must adhere to. It provides methods for initializing the service and
// retrieving its name.
type IsRuntimeService interface {
	// Init initializes the service and establishes a connection to the
	// service IsRuntime.
	Init(rt IsRuntime)

	// Name returns the name of the service.
	Name() string
}

// HasDependencies represents a service that can resolve its dependencies.
//
// The HasDependencies interface extends the Service interface and adds
// methods for managing dependencies. It allows services to declare whether
// their dependencies are resolved, as well as a method that attempts to resolve
// those dependencies with the given runtime.
//
// The Runtime will use this interface automatically when a service is added.
// You do not need to implement this interface, it is optional. You would want
// to do this when you have services that depend upon each other to operate
type HasDependencies interface {
	IsRuntimeService

	// DependenciesResolved returns true if all dependencies are resolved. This
	// is up to the service.
	DependenciesResolved() bool

	// ResolveDependencies attempts to resolve the dependencies of the
	// service using the provided IsRuntime.
	ResolveDependencies(IsRuntime)
}

// HasLogger is an interface for components that require a logger instance.
//
// The HasLogger interface represents components that depend on a logger for
// logging purposes. It defines a method to set the logger instance.
type HasLogger interface {
	IsRuntimeService
	// UseLogger sets the logger instance for the component.
	BindLogger(logger *zerolog.Logger)
	// Logger yields the logger instance for the component.
	Logger() *zerolog.Logger
}

// HasGracefulShutdown is an interface for services that require graceful shutdown handling.
//
// The HasGracefulShutdown interface extends the IsRuntimeService interface and adds
// a method for performing custom actions during graceful shutdown.
type HasGracefulShutdown interface {
	IsRuntimeService

	// OnShutdown is called during the graceful shutdown process to perform
	// custom actions before the service is stopped.
	OnShutdown()
}

// EventHandlerServiceAdded is an optional interface. If implemented, it will automatically bind to the
// "Service Added" runtime event, allowing the object to respond when a new service is added.
type EventHandlerServiceAdded interface {
	OnServiceAdded(args ...interface{})
}

// EventHandlerServiceRemoved is an optional interface. If implemented, it will automatically bind to the
// "Service Removed" runtime event, enabling the implementor to respond when a service is removed.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerServiceRemoved interface {
	OnServiceRemoved(args ...interface{})
}

// EventHandlerServiceInitialized is an optional interface. If implemented, it will automatically bind to the
// "Service Initialized" runtime event, enabling the implementor to respond when a service is initialized.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerServiceInitialized interface {
	OnServiceInitialized(args ...interface{})
}

// EventHandlerServiceEventsBound is an optional interface. If implemented, it will automatically bind to the
// "Service Events Bound" runtime event, enabling the implementor to respond when events are bound to a service.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerServiceEventsBound interface {
	OnServiceEventsBound(args ...interface{})
}

// EventHandlerServiceLoggerBound is an optional interface. If implemented, it will automatically bind to the
// "Service Logger Bound" runtime event, enabling the implementor to respond when a logger is bound to a service.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerServiceLoggerBound interface {
	OnServiceLoggerBound(args ...interface{})
}

// EventHandlerRuntimeRunLoopInitiated is an optional interface. If implemented, it will automatically bind to the
// "Runtime Run Loop Initiated" runtime event, enabling the implementor to respond when the runtime run loop is initiated.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerRuntimeRunLoopInitiated interface {
	OnRuntimeRunLoopInitiated(args ...interface{})
}

// EventHandlerRuntimeShutdownInitiated is an optional interface. If implemented, it will automatically bind to the
// "Runtime Shutdown Initiated" runtime event, enabling the implementor to respond when the runtime is preparing to shut down.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerRuntimeShutdownInitiated interface {
	OnRuntimeShutdownInitiated(args ...interface{})
}

// EventHandlerDependencyResolutionStarted is an optional interface. If implemented, it will automatically bind to the
// "Dependency Resolution Started" runtime event, enabling the implementor to respond when dependency resolution starts.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerDependencyResolutionStarted interface {
	OnDependencyResolutionStarted(args ...interface{})
}

// EventHandlerDependencyResolutionEnded is an optional interface. If implemented, it will automatically bind to the
// "Dependency Resolution Ended" runtime event, enabling the implementor to respond when dependency resolution ends.
// When the event is emitted, the declared method will be called and passed the arguments from the emitter.
type EventHandlerDependencyResolutionEnded interface {
	OnDependencyResolutionEnded(args ...interface{})
}
