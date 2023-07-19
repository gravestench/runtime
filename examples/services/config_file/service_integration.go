package config_file

import (
	"github.com/gravestench/runtime"
)

// this is a static check that my service satisfies the
// recipe below. This should prevent the code from compiling
// if this service should not implement these interfaces.
var _ recipe = &Service{}

type recipe interface {
	runtime.IsRuntimeService
	runtime.HasLogger
	runtime.UsesEventBus
	Manager
}

// The following interfaces are to be used much like the service interfaces
// found inside of runtime/pkg. These can be used by other services to
// declare and resolve their dependencies to the service defined in this
// package.

// HasConfig represents a something with a configuration file path and retrieval methods.
type HasConfig interface {
	ConfigFilePath() string   // ConfigFilePath returns the path to the configuration file.
	Config() (*Config, error) // Config retrieves the configuration from the file.
}

// HasDefaultConfig represents something with a default configuration.
type HasDefaultConfig interface {
	HasConfig
	DefaultConfig() Config // DefaultConfig returns the default configuration.
}

// Manager represents something that manages configurations.
type Manager interface {
	GetPath(path string) string
	ConfigDirectory() string                   // ConfigDirectory returns the directory path where configurations are stored.
	SetConfigDirectory(string) error           // SetConfigDirectory sets the directory path for configurations.
	Configs() map[string]*Config               // Configs returns all configurations stored in the service.
	GetConfig(string) (*Config, error)         // GetConfig retrieves a configuration by its path.
	CreateConfig(path string) (*Config, error) // CreateConfig creates a new configuration file at the specified path.
	LoadConfig(string) (*Config, error)        // LoadConfig loads a configuration from the specified path.
	SaveConfig(string) error                   // SaveConfig saves a configuration to the specified path.
}
