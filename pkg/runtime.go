package pkg

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	ee "github.com/gravestench/eventemitter"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/pkg/events"
)

var _ IsRuntime = &Runtime{}

// Runtime represents a collection of runtime services.
type Runtime struct {
	name     string
	quit     chan os.Signal
	services []IsRuntimeService
	logger   *zerolog.Logger
	stdOut   io.Writer
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
		stdOut: os.Stdout,
	}

	// the runtime itself is a service that binds handlers to its own events
	r.Add(r)

	return r
}

func (r *Runtime) Init(_ IsRuntime) {
	if r.services != nil {
		return
	}

	r.logger = r.newLogger(r, zerolog.InfoLevel)

	r.logger.Info().Msgf("initializing")

	r.quit = make(chan os.Signal, 1)
	signal.Notify(r.quit, os.Interrupt)

	r.services = make([]IsRuntimeService, 0)
}

// Add a single service to the Runtime manager.
func (r *Runtime) Add(service IsRuntimeService) *sync.WaitGroup {
	r.Init(nil) // always ensure runtime is init
	r.bindEventHandlerInterfaces(service)

	var wg sync.WaitGroup

	if service != r {
		r.logger.Info().Msgf("preparing service %q", service.Name())
	}

	// Check if the service uses a logger
	if loggerUser, ok := service.(HasLogger); ok {
		wg.Add(1)
		loggerUser.BindLogger(r.newLogger(service, r.logger.GetLevel()))
		r.events.Emit(events.EventServiceLoggerBound, service).Wait()
		wg.Done()
	}

	r.services = append(r.services, service)

	// Check if the service is a HasDependencies
	if resolver, ok := service.(HasDependencies); ok {
		// Resolve dependencies before initialization
		go func() {
			wg.Add(1)
			r.resolveDependenciesAndInit(resolver)
			r.events.Emit(events.EventServiceAdded, service)
			wg.Done()
		}()
	} else {
		// No dependencies to resolve, directly initialize the service
		go func() {
			wg.Add(1)
			r.initService(service)
			r.events.Emit(events.EventServiceAdded, service)
			wg.Done()
		}()
	}

	return &wg
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
		r.newLogger(service, r.logger.GetLevel()).Debug().Msgf("initializing")
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
func (r *Runtime) Remove(service IsRuntimeService) *sync.WaitGroup {
	wg := r.events.Emit(events.EventServiceRemoved)

	for i, svc := range r.services {
		if svc == service {
			r.logger.Info().Msgf("removing %q service", service.Name())
			r.services = append(r.services[:i], r.services[i+1:]...)
			break
		}
	}

	return wg
}

// Shutdown sends an interrupt signal to the Runtime, indicating it should exit.
func (r *Runtime) Shutdown() *sync.WaitGroup {
	wg := r.events.Emit(events.EventRuntimeShutdownInitiated)

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

	return wg
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

	r.Shutdown().Wait()
	os.Exit(0)
}

// Events yields the global event bus for the runtime
func (r *Runtime) Events() *ee.EventEmitter {
	return r.events
}

func (r *Runtime) bindEventHandlerInterfaces(service IsRuntimeService) {
	if handler, ok := service.(EventHandlerServiceAdded); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventServiceAdded' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventServiceAdded, handler.OnServiceAdded)
	}

	if handler, ok := service.(EventHandlerServiceRemoved); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventServiceRemoved' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventServiceRemoved, handler.OnServiceRemoved)
	}

	if handler, ok := service.(EventHandlerServiceInitialized); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventServiceInitialized' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventServiceInitialized, handler.OnServiceInitialized)
	}

	if handler, ok := service.(EventHandlerServiceEventsBound); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventServiceEventsBound' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventServiceEventsBound, handler.OnServiceEventsBound)
	}

	if handler, ok := service.(EventHandlerServiceLoggerBound); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventServiceLoggerBound' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventServiceLoggerBound, handler.OnServiceLoggerBound)
	}

	if handler, ok := service.(EventHandlerRuntimeRunLoopInitiated); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventRuntimeRunLoopInitiated' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventRuntimeRunLoopInitiated, handler.OnRuntimeRunLoopInitiated)
	}

	if handler, ok := service.(EventHandlerRuntimeShutdownInitiated); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventRuntimeShutdownInitiated' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventRuntimeShutdownInitiated, handler.OnRuntimeShutdownInitiated)
	}

	if handler, ok := service.(EventHandlerDependencyResolutionStarted); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventDependencyResolutionStarted' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventDependencyResolutionStarted, handler.OnDependencyResolutionStarted)
	}

	if handler, ok := service.(EventHandlerDependencyResolutionEnded); ok {
		if service != r {
			r.logger.Info().Msgf("bound 'EventDependencyResolutionEnded' event handler for service %q", service.Name())
		}
		r.Events().On(events.EventDependencyResolutionEnded, handler.OnDependencyResolutionEnded)
	}
}

func (r *Runtime) OnServiceAdded(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		if service != r {
			r.logger.Info().Msgf("service %q has been added", service.Name())
		}
	}
}

func (r *Runtime) OnRuntimeShutdownInitiated(_ ...any) {
	r.logger.Warn().Msg("initiating graceful shutdown")
}

func (r *Runtime) OnServiceRemoved(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("removed service %q", service.Name())
	}
}

func (r *Runtime) OnServiceInitialized(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("service %q initialized", service.Name())
	}
}

func (r *Runtime) OnServiceEventsBound(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("events bound for service %q", service.Name())
	}
}

func (r *Runtime) OnServiceLoggerBound(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("logger bound for service %q", service.Name())
	}
}

func (r *Runtime) OnRuntimeRunLoopInitiated(_ ...any) {
	r.logger.Debug().Msg("run loop started")
}

func (r *Runtime) OnDependencyResolutionStarted(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("dependency resolution started for service %q", service.Name())
	}
}

func (r *Runtime) OnDependencyResolutionEnded(args ...any) {
	if len(args) < 1 {
		return
	}

	if service, ok := args[0].(IsRuntimeService); ok {
		r.logger.Debug().Msgf("dependency resolution completed for service %q", service.Name())
	}
}
