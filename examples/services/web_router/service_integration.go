package web_router

import (
	"github.com/gin-gonic/gin"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
)

var (
	_ runtime.Service              = &Service{}
	_ runtime.HasLogger            = &Service{}
	_ runtime.HasDependencies      = &Service{}
	_ config_file.HasDefaultConfig = &Service{}
	_ IsWebRouter                  = &Service{}
)

// Router is just responsible for yielding the root route handler.
// Services will use this in order to set up their own routes.
type IsWebRouter interface {
	RouteRoot() *gin.Engine
	Reload()
}

// HasRouteSlug describes a service that has an identifier that is used
// as a prefix for its subroutes
type HasRouteSlug interface {
	Slug() string
}

// IsRouteInitializer is a type of service that will
// set up its own web routes using a base route group
type IsRouteInitializer interface {
	runtime.Service
	InitRoutes(*gin.RouterGroup)
}
