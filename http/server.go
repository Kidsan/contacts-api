package http

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	contactsapi "github.com/kidsan/contacts-api"
	pb "github.com/kidsan/contacts-api/protobuffer"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type HTTPServer struct {
	logger     *zap.Logger
	router     *chi.Mux
	connection *gorm.DB
	config     contactsapi.ServerConfig
}

type GRPCServer struct {
	logger     *zap.Logger
	server     *grpc.Server
	connection *gorm.DB
	config     contactsapi.ServerConfig
}

func NewHTTPServer(config contactsapi.Config, logger *zap.Logger, connection *gorm.DB) *HTTPServer {
	r := chi.NewRouter()
	s := &HTTPServer{
		config:     config.Server,
		logger:     logger,
		connection: connection,
		router:     r,
	}

	s.router.Get("/metrics", s.handleMetrics())
	s.router.Route("/api/contacts", s.buildContactRouter())
	return s
}

func (s *HTTPServer) Start() {
	s.logger.Info(fmt.Sprintf("HTTP Application listening on port %d", s.config.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router)
}

func NewGRPCServer(config contactsapi.Config, logger *zap.Logger, connection *gorm.DB) *GRPCServer {
	grpcServer := &GRPCServer{
		config:     config.Server,
		logger:     logger,
		connection: connection,
		server:     grpc.NewServer(),
	}
	pb.RegisterContactsServer(grpcServer.server, grpcServer.buildContactServer())
	return grpcServer
}

func (s *GRPCServer) Start() {
	s.logger.Info(fmt.Sprintf("GRPC Application listening on port %d", s.config.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.config.Port))
	if err != nil {
		s.logger.Sugar().Fatalf("failed to listen: %v", err)
	}
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
