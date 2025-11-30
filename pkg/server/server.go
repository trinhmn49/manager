package server

import (
	"fmt"
	"manager/pkg/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config config.AppConfig
}

type RouteRegister interface {
	RegisterRoutes(r *gin.RouterGroup)
}

func New(config config.AppConfig) *Server {
	var r *gin.Engine
	if config.RunMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r = gin.Default()
	return &Server{
		router: r,
		config: config,
	}
}

func (s *Server) SetupRoute(routes []RouteRegister) {
	apiVersion := s.router.Group("/v1")
	for _, route := range routes {
		route.RegisterRoutes(apiVersion)
	}
}

func (s *Server) Start() error {
	return s.router.Run(fmt.Sprintf(":%d", s.config.Port))
}
