package grpc

import (
	"fmt"
	contextcomponent "github.com/Alp4ka/classifier-aaS/internal/components/context"
	contextrepository "github.com/Alp4ka/classifier-aaS/internal/components/context/repository"
	schemacomponent "github.com/Alp4ka/classifier-aaS/internal/components/schema"
	schemarepository "github.com/Alp4ka/classifier-aaS/internal/components/schema/repository"
	grpcpkg "github.com/Alp4ka/classifier-aaS/pkg/grpc"
	api "github.com/Alp4ka/classifier-api"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	api.UnimplementedGWManagerServiceServer

	cfg            Config
	grpcServer     *grpc.Server
	contextService contextcomponent.Service
}

func New(cfg Config) *Server {
	return &Server{
		cfg: cfg,
		contextService: contextcomponent.NewService(
			contextcomponent.Config{
				SchemaService: schemacomponent.NewService(
					schemacomponent.Config{
						Repository: schemarepository.NewRepository(cfg.DB),
					},
				),
				Repository: contextrepository.NewRepository(cfg.DB),
			},
		),
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmiddleware.ChainStreamServer(grpcrecovery.StreamServerInterceptor(), grpcpkg.ServerStreamInterceptor()),
		),
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(grpcrecovery.UnaryServerInterceptor(), grpcpkg.UnaryServerInterceptor()),
		),
	)
	reflection.Register(s.grpcServer)
	api.RegisterGWManagerServiceServer(s.grpcServer, s)
	mlogger.L().Info("Listening GRPC API server", field.Int("port", s.cfg.Port))
	return s.grpcServer.Serve(lis)
}

func (s *Server) Close() error {
	s.grpcServer.Stop()
	return nil
}
