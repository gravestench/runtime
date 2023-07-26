package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	ee "github.com/gravestench/eventemitter"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/pkg/events"
)

var _ IsRuntime = &Runtime{}

// Runtime represents a manager for runtime services.
type Runtime struct {
	name     string
	quit     chan os.Signal
	services []IsRuntimeService
	logger   *zerolog.Logger
	events   *ee.EventEmitter
}

// New creates a new instance of a Runtime.
func New(args ...string) *Runtime {
	name := "Runtime"

	if len(args) > 0 {
		name = strings.Join(args, " ")
	}

	r := &Runtime{
		name:   name,
		events: ee.New(),
	}

	for event, handler := range map[string]func(...any){
		events.EventServiceAdded:                r.onEventServiceAdded,
		events.EventServiceRemoved:              r.onEventServiceRemoved,
		events.EventServiceInitialized:          r.onEventServiceInitialized,
		events.EventServiceEventsBound:          r.onEventServiceEventsBound,
		events.EventServiceLoggerBound:          r.onEventServiceLoggerBound,
		events.EventRuntimeShutdownInitiated:    r.onEventRuntimeShutdownInitiated,
		events.EventRuntimeRunLoopInitiated:     r.onEventRuntimeRunLoopInitiated,
		events.EventDependencyResolutionStarted: r.onEventDependencyResolutionStarted,
		events.EventDependencyResolutionEnded:   r.onEventDependencyResolutionEnded,
	} {
		r.events.On(event, handler)
	}

	r.ensureInit()
	return r
}

// ensureInit initializes the Runtime manager.
func (r *Runtime) ensureInit() {
	if r.services != nil {
		return
	}

	r.logger = newLogger(r, zerolog.InfoLevel)

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
		loggerUser.BindLogger(newLogger(service, r.logger.GetLevel()))
		r.events.Emit(events.EventServiceLoggerBound, service)
	}

	r.services = append(r.services, service)
	r.events.Emit(events.EventServiceAdded, service)

	// Check if the service is a HasDependencies
	if resolver, ok := service.(HasDependencies); ok {
		// Resolve dependencies before initialization
		go r.resolveDependenciesAndInit(resolver)
	} else {
		// No dependencies to resolve, directly initialize the service
		go r.initService(service)
	}
}

func (r *Runtime) SetLogLevel(level zerolog.Level) {
	r.logger.Info().Msgf("setting log level to %s", level)

	// set the log-level for the runtime's logger
	instance := r.logger.Level(level)
	r.logger = &instance

	// set the log level for each service that has a logger
	for _, service := range r.Services() {
		candidate, ok := service.(HasLogger)
		if !ok {
			continue
		}

		candidateLogger := candidate.Logger().Level(level)
		candidate.BindLogger(&candidateLogger)
	}
}

func (r *Runtime) resolveDependenciesAndInit(resolver HasDependencies) {
	r.events.Emit(events.EventDependencyResolutionStarted, resolver)

	// Check if all dependencies are resolved
	for !resolver.DependenciesResolved() {
		resolver.ResolveDependencies(r)
		time.Sleep(time.Millisecond * 10)
	}

	r.events.Emit(events.EventDependencyResolutionEnded, resolver)

	// All dependencies resolved, initialize the service
	r.initService(resolver)
}

// initService initializes a service and adds it to the Runtime manager.
func (r *Runtime) initService(service IsRuntimeService) {
	if l, ok := service.(HasLogger); ok && l.Logger() != nil {
		l.Logger().Debug().Msg("initializing")
	} else {
		newLogger(service, r.logger.GetLevel()).Debug().Msgf("initializing")
	}

	if candidate, ok := service.(UsesEventBus); ok {
		candidate.BindsEvents(r.events)
		r.events.Emit(events.EventServiceEventsBound)
	}

	// Initialize the service
	service.Init(r)

	r.events.Emit(events.EventServiceInitialized, service)
}

// Services returns a pointer to a slice of interfaces representing the services managed by the Runtime.
func (r *Runtime) Services() []IsRuntimeService {
	duplicate := append([]IsRuntimeService{}, r.services...)
	return duplicate
}

// Remove a specific service from the Runtime manager.
func (r *Runtime) Remove(service IsRuntimeService) {
	r.events.Emit(events.EventServiceRemoved)
	for i, svc := range r.services {
		if svc == service {
			r.logger.Info().Msgf("removing %q service", service.Name())
			r.services = append(r.services[:i], r.services[i+1:]...)
			break
		}
	}
}

// Shutdown sends an interrupt signal to the Runtime, indicating it should exit.
func (r *Runtime) Shutdown() {
	r.events.Emit(events.EventRuntimeShutdownInitiated)

	for _, service := range r.services {
		if quitter, ok := service.(HasGracefulShutdown); ok {

			if l, ok := quitter.(HasLogger); ok && l.Logger() != nil {
				l.Logger().Info().Msg("shutting down")
			} else {
				r.logger.Info().Msgf("shutting down %q service", service.Name())
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
	r.events.Emit(events.EventRuntimeRunLoopInitiated)

	<-r.quit              // blocks until signal is recieved
	fmt.Printf("\033[2D") // Remove ^C from stdout

	r.Shutdown()
}

func (r *Runtime) onEventServiceAdded(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Info().Msgf("adding service %q", service.Name())
	}
}

func (r *Runtime) onEventRuntimeShutdownInitiated(_ ...any) {
	r.logger.Warn().Msg("initiating graceful shutdown")
}

func (r *Runtime) onEventServiceRemoved(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("removing service %q", service.Name())
	}
}
func (r *Runtime) onEventServiceInitialized(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("service %q initialization complete", service.Name())
	}
}
func (r *Runtime) onEventServiceEventsBound(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("events bound for service %q", service.Name())
	}
}
func (r *Runtime) onEventServiceLoggerBound(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("logger bound for service %q", service.Name())
	}
}
func (r *Runtime) onEventRuntimeRunLoopInitiated(_ ...any) {
	r.logger.Debug().Msg("beginning run loop")
}

func (r *Runtime) onEventDependencyResolutionStarted(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("dependency resolution started for service %q", service.Name())
	}
}

func (r *Runtime) onEventDependencyResolutionEnded(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("dependency resolution completed for service %q", service.Name())
	}
}
