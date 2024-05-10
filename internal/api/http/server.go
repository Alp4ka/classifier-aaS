package http

import (
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
	globaltelemtry "github.com/Alp4ka/classifier-aaS/internal/telemetry"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

const (
	_appName = "http-api"
)

type Server struct {
	app           *fiber.App
	schemaService schemacomponent.Service
	port          int
	rateLimit     int
}

func NewHTTPServer(cfg Config) *Server {
	server := &Server{
		schemaService: cfg.SchemaService,
		port:          cfg.Port,
		rateLimit:     cfg.RateLimit,
	}
	server.app = fiber.New(
		fiber.Config{
			AppName:               _appName,
			ErrorHandler:          server.mwErrorHandler,
			DisableStartupMessage: true,
			DisableKeepalive:      true,
		},
	)

	return server
}

func (s *Server) configureRouting() {
	// Middlewares.
	s.app.Use(s.mwRecoverer())
	s.app.Use(s.mwCors())
	s.app.Use(s.mwSwagger())
	s.app.Use(s.mwRequestID())
	s.app.Use(s.mwLogging())

	// API Group
	apiGroup := s.app.Group("/api")
	apiGroup.Use(s.mwRateLimiter())
	apiGroup.Use(s.mwContentChecker())

	// SchemaReq.
	schemaGroup := apiGroup.Group("/schema")
	schemaGroup.Get("/:id", s.hGetActualSchema)
	schemaGroup.Post("", s.hCreateSchema)
	schemaGroup.Put("", s.hUpdateSchema)
}

func (s *Server) Run() error {
	s.configureRouting()
	mlogger.L().Info("Listening HTTP API server", field.Int("port", s.port))
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}

func (s *Server) WithMetrics() *Server {
	prometheus := fiberprometheus.New(globaltelemtry.Namespace)
	prometheus.RegisterAt(s.app, "/metrics")
	return s
}

func (s *Server) Close() error {
	return s.app.Shutdown()
}
