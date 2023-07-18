package config_file

import (
	ee "github.com/gravestench/eventemitter"
)

const (
	EventConfigChanged = "config file changed"
)

func (s *Service) BindsEvents(emitter *ee.EventEmitter) {
	s.events = emitter
}
