package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/web_router"
	"github.com/gravestench/runtime/pkg"
)

var (
	_ runtime.Service               = &Service{}
	_ web_router.IsRouteInitializer = &Service{}
)

type Service struct{}

func (s *Service) Init(rt pkg.IsRuntime) {

}

func (s *Service) Name() string {
	return "Example Service With Web Routes"
}

func (s *Service) InitRoutes(group *gin.RouterGroup) {
	group.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "It hella works!")
	})
}
