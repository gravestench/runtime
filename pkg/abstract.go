package pkg

import (
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

	Shutdown()
}

// IsRuntimeService represents a generic service within a runtime.
//
// The IsRuntimeService interface defines the contract that all services in the
// IsRuntime must adhere to. It provides methods for initializing the service and
// retrieving its name.
type IsRuntimeService interface {
	// Init initializes the service and establishes a connection to the
	// service IsRuntime.
	Init(IsRuntime)

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
