package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var _ IsRuntime = &Runtime{}

// Runtime represents a manager for runtime services.
type Runtime struct {
	name     string
	quit     chan os.Signal
	services []IsRuntimeService
	logger   *zerolog.Logger
}

// New creates a new instance of a Runtime.
func New(args ...string) *Runtime {
	name := "Runtime"

	if len(args) > 0 {
		name = strings.Join(args, " ")
	}

	r := &Runtime{
		name: name,
	}

	r.ensureInit()
	return r
}

// ensureInit initializes the Runtime manager.
func (r *Runtime) ensureInit() {
	if r.services != nil {
		return
	}

	r.logger = newLogger(r)

	r.logger.Info().Msgf("initializing")

	r.quit = make(chan os.Signal, 1)
	signal.Notify(r.quit, os.Interrupt)

	r.services = make([]IsRuntimeService, 0)
}

// Add a single service to the Runtime manager.
func (r *Runtime) Add(service IsRuntimeService) {
	r.ensureInit()

	// Check if the service uses a logger
	if loggerUser, ok := service.(HasLogger); ok {
		loggerUser.BindLogger(newLogger(service))
	}

	r.services = append(r.services, service)

	// Check if the service is a HasDependencies
	if resolver, ok := service.(HasDependencies); ok {
		// Resolve dependencies before initialization
		go r.resolveDependenciesAndInit(resolver)
	} else {
		// No dependencies to resolve, directly initialize the service
		go r.initService(service)
	}
}

func (r *Runtime) resolveDependenciesAndInit(resolver HasDependencies) {
	r.logger.Info().Msgf("resolving dependencies for '%s'", resolver.Name())

	// Check if all dependencies are resolved
	for !resolver.DependenciesResolved() {
		resolver.ResolveDependencies(r)
		time.Sleep(time.Millisecond * 10)
	}

	r.logger.Info().Msgf("dependencies resolved for '%s'", resolver.Name())

	// All dependencies resolved, initialize the service
	r.initService(resolver)
}

// initService initializes a service and adds it to the Runtime manager.
func (r *Runtime) initService(service IsRuntimeService) {
	if l, ok := service.(HasLogger); ok && l.Logger() != nil {
		l.Logger().Info().Msg("initializing")
	} else {
		r.logger.Info().Msgf("initializing '%s' service", service.Name())
	}

	// Initialize the service
	service.Init(r)
}

// Services returns a pointer to a slice of interfaces representing the services managed by the Runtime.
func (r *Runtime) Services() []IsRuntimeService {
	duplicate := append([]IsRuntimeService{}, r.services...)
	return duplicate
}

// Remove a specific service from the Runtime manager.
func (r *Runtime) Remove(service IsRuntimeService) {
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
	r.logger.Info().Msg("initiating graceful shutdown")

	for _, service := range r.services {
		if quitter, ok := service.(HasGracefulShutdown); ok {

			if l, ok := quitter.(HasLogger); ok && l.Logger() != nil {
				l.Logger().Info().Msg("shutting down")
			} else {
				r.logger.Info().Msgf("shutting down '%s' service", service.Name())
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
