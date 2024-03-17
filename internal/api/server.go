package api

import (
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
)

const (
	_appName     = "http-api"
	_readTimeout = 5 * time.Second
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
			AppName:      _appName,
			ErrorHandler: server.mwErrorHandler,
			ReadTimeout:  _readTimeout,
		},
	)

	return server
}

func (s *HTTPServer) configureRouting() {
	// Middlewares.
	s.app.Use(requestid.New())
	s.app.Use(s.mwGetRecoverer())
	s.app.Use(s.mwLogging)

	// API Group
	apiGroup := s.app.Group("/api")
	apiGroup.Use(s.mwGetRateLimiter(s.rateLimit))
	apiGroup.Use(s.mwContentChecker)

	// Schema.
	schemaGroup := apiGroup.Group("/schema")
	schemaGroup.Post("", s.hCreateSchema)
	schemaGroup.Patch("", s.hPatchSchema)
	schemaGroup.Get("/:id", s.hGetSchema)
}

func (s *HTTPServer) Run() error {
	s.configureRouting()
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}

func (s *HTTPServer) Close() error {
	return s.app.Shutdown()
}
