package api

import (
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
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

type HTTPServer struct {
	app            *fiber.App
	schemaService  schemacomponent.Service
	contextService contextcomponent.Service
	port           int
	rateLimit      int
}

func NewHTTPServer(cfg Config) *HTTPServer {
	server := &HTTPServer{
		contextService: cfg.ContextService,
		schemaService:  cfg.SchemaService,
		port:           cfg.Port,
		rateLimit:      cfg.RateLimit,
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

func (s *HTTPServer) configureRouting() {
	// Middlewares.
	s.app.Use(s.mwGetRecoverer())
	s.app.Use(s.mwGetSwagger())
	s.app.Use(s.mwGetRequestIDer())
	s.app.Use(s.mwLogging())

	// API Group
	apiGroup := s.app.Group("/api")
	apiGroup.Use(s.mwGetRateLimiter())
	apiGroup.Use(s.mwContentChecker())

	// Schema.
	schemaGroup := apiGroup.Group("/schema")
	schemaGroup.Get("/:id", s.hGetActualSchema)
	schemaGroup.Post("", s.hCreateSchema)
	schemaGroup.Put("", s.hUpdateSchema)
}

func (s *HTTPServer) Run() error {
	s.configureRouting()
	mlogger.L().Info("Listening API server", field.Int("port", s.port))
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}

func (s *HTTPServer) WithMetrics() *HTTPServer {
	prometheus := fiberprometheus.New(globaltelemtry.Namespace)
	prometheus.RegisterAt(s.app, "/metrics")
	return s
}

func (s *HTTPServer) Close() error {
	return s.app.Shutdown()
}
