package pkg

import (
	"github.com/rs/zerolog"
)

// RuntimeInterface is responsible for managing services within the system.
//
// The RuntimeInterface interface defines the operations that can be performed on
// services, such as adding, removing, and retrieving services. It acts as a
// container for services and provides methods for managing their lifecycle.
type RuntimeInterface interface {
	// Add a single service to the RuntimeInterface.
	Add(RuntimeServiceInterface)

	// Remove a specific service from the RuntimeInterface.
	Remove(RuntimeServiceInterface)

	// Services returns a pointer to a slice of interfaces representing the
	// services currently managed by the service RuntimeInterface.
	Services() *[]interface{}
}

// RuntimeServiceInterface represents a generic service within the RuntimeInterface.
//
// The RuntimeServiceInterface interface defines the contract that all services in the
// RuntimeInterface must adhere to. It provides methods for initializing the service and
// retrieving its name.
type RuntimeServiceInterface interface {
	// Init initializes the service and establishes a connection to the
	// service RuntimeInterface.
	Init(RuntimeInterface)

	// Name returns the name of the service.
	Name() string
}

// DependencyResolver represents a service that can resolve its dependencies.
//
// The DependencyResolver interface extends the Service interface and adds
// methods for managing dependencies. It allows services to declare their
// dependencies, resolve them, and initialize themselves once all dependencies
// are satisfied.
type DependencyResolver interface {
	// RuntimeServiceInterface represents the generic service interface.
	RuntimeServiceInterface

	// DependenciesResolved returns true if all dependencies are resolved. This
	// is up to the service.
	DependenciesResolved() bool

	// ResolveDependencies attempts to resolve the dependencies of the
	// service using the provided RuntimeInterface.
	ResolveDependencies(RuntimeInterface)
}

// UsesLogger is an interface for components that require a logger instance.
//
// The UsesLogger interface represents components that depend on a logger for
// logging purposes. It defines a method to set the logger instance.
type UsesLogger interface {
	// UseLogger sets the logger instance for the component.
	UseLogger(logger *zerolog.Logger)
	// Logger yields the logger instance for the component.
	Logger() *zerolog.Logger
}
