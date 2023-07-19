package desktop_notification

import (
	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

// Ensure that Service implements the required interfaces.
var (
	_ runtime.Service              = &Service{}
	_ runtime.HasLogger            = &Service{}
	_ runtime.HasDependencies      = &Service{}
	_ config_file.HasDefaultConfig = &Service{}
	_ SendsNotifications           = &Service{}
)

type SendsNotifications interface {
	Notify(title, message, appIcon string)
	Alert(title, message, appIcon string)
}
