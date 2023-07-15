package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var _ RuntimeInterface = &Runtime{}

// Runtime represents a manager for RuntimeInterface services.
type Runtime struct {
	name     string
	quit     chan os.Signal
	services []RuntimeServiceInterface
	logger   *zerolog.Logger
}

// NewRuntime creates a new instance of the Runtime manager.
func New(args ...string) *Runtime {
	name := "Runtime"

	if len(args) > 0 {
		name = strings.Join(args, " ")
	}

	r := &Runtime{
		name: name,
	}

	r.init()
	return r
}

// init initializes the Runtime manager.
func (r *Runtime) init() {
	if r.services != nil {
		return
	}

	r.logger = newLogger(r.Name())

	r.logger.Info().Msgf("initializing")

	r.quit = make(chan os.Signal, 1)
	signal.Notify(r.quit, os.Interrupt)

	r.services = make([]RuntimeServiceInterface, 0)

	// Prevent deadlock panic by continuously sleeping
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

// Add a single service to the Runtime manager.
func (r *Runtime) Add(service RuntimeServiceInterface) {
	r.init()

	// Check if the service uses a logger
	if loggerUser, ok := service.(UsesLogger); ok {
		loggerUser.UseLogger(newLogger(service.Name()))
	}

	// Check if the service is a DependencyResolver
	if resolver, ok := service.(DependencyResolver); ok {
		// Resolve dependencies before initialization
		r.resolveDependenciesAndInit(resolver)
	} else {
		// No dependencies to resolve, directly initialize the service
		r.initService(service)
	}
}

// resolveDependenciesAndInit resolves dependencies for a DependencyResolver service.
func (r *Runtime) resolveDependenciesAndInit(resolver DependencyResolver) {
	if l, ok := resolver.(UsesLogger); ok && l.Logger() != nil {
		l.Logger().Info().Msg("resolving dependencies")
	} else {
		r.logger.Info().Msgf("resolving dependencies for '%s'", resolver.Name())
	}

	// Attempt to resolve dependencies
	resolver.ResolveDependencies(r)

	// Check if all dependencies are resolved
	for !resolver.DependenciesResolved() {
		r.resolveDependenciesAndInit(resolver)
		time.Sleep(time.Millisecond * 10)
	}

	// All dependencies resolved, initialize the service
	r.initService(resolver)
}

// initService initializes a service and adds it to the Runtime manager.
func (r *Runtime) initService(service RuntimeServiceInterface) {
	if l, ok := service.(UsesLogger); ok && l.Logger() != nil {
		l.Logger().Info().Msg("initializing")
	} else {
		r.logger.Info().Msgf("initializing '%s' service", service.Name())
	}

	r.services = append(r.services, service)

	// Initialize the service
	service.Init(r)
}

// Services returns a pointer to a slice of interfaces representing the services managed by the Runtime.
func (r *Runtime) Services() *[]RuntimeServiceInterface {
	duplicate := append([]RuntimeServiceInterface{}, r.services...)
	return &duplicate
}

// Remove a specific service from the Runtime manager.
func (r *Runtime) Remove(service RuntimeServiceInterface) {
	for i, svc := range r.services {
		if svc == service {
			r.logger.Info().Msgf("removing '%s' service", service.Name())
			r.services = append(r.services[:i], r.services[i+1:]...)
			break
		}
	}
}

// Shutdown sends an interrupt signal to the Runtime, indicating it should exit.
func (r *Runtime) Shutdown() {
	r.quit <- os.Interrupt
	r.logger.Warn().Msg("initiating graceful shutdown")

	for _, service := range r.services {
		if quitter, ok := service.(HasGracefulShutdown); ok {

			if l, ok := quitter.(UsesLogger); ok && l.Logger() != nil {
				l.Logger().Warn().Msg("shutting down")
			} else {
				r.logger.Warn().Msgf("shutting down '%s' service", service.Name())
			}

			quitter.OnShutdown()
		}
	}

	r.logger.Info().Msg("exiting")
}

// Name returns the name of the Runtime manager.
func (r *Runtime) Name() string {
	return r.name
}

// Run starts the Runtime manager and waits for an interrupt signal to exit.
func (r *Runtime) Run() {
	r.logger.Info().Msg("beginning run loop")

	<-r.quit              // blocks until signal is recieved
	fmt.Printf("\033[2D") // Remove ^C from stdout

	r.Shutdown()
}
