package grpc

import (
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/components/context"
	api "github.com/Alp4ka/classifier-api"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	mapset "github.com/deckarep/golang-set/v2"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	api.UnimplementedGWManagerServiceServer

	contextService context.Service
	grpcServer     *grpc.Server

	ridStorage mapset.Set[string]

	port int
}

func New(cfg Config) *Server {
	return &Server{
		contextService: cfg.ContextService,
		port:           cfg.Port,
		ridStorage:     mapset.NewSet[string](),
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(grpc_recovery.StreamServerInterceptor()),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(grpc_recovery.UnaryServerInterceptor()),
		),
	)
	reflection.Register(s.grpcServer)
	api.RegisterGWManagerServiceServer(s.grpcServer, s)
	mlogger.L().Info("Listening GRPC API server", field.Int("port", s.port))
	return s.grpcServer.Serve(lis)
}

func (s *Server) Close() error {
	s.grpcServer.Stop()
	return nil
}

func (s *Server) IsDuplicate(req *api.ProcessRequest) bool {
	return s.ridStorage.Contains(req.GetRequestId())
}

func (s *Server) StoreRequest(req *api.ProcessRequest) {
	s.ridStorage.Add(req.GetRequestId())
}
