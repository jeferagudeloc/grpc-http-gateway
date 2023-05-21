package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/grpc/order"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/grpc/user"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
)

type grpcServer struct {
	log logger.Logger
}

func NewGRPCServer(
	log logger.Logger,
) *grpcServer {
	return &grpcServer{
		log: log,
	}
}

func (s *grpcServer) Listen() {
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	// main server
	grpcServer := grpc.NewServer()

	// server declarations
	orderServer := order.Server{}
	userServer := user.Server{}

	order.RegisterOrderHandlerServer(grpcServer, &orderServer)
	user.RegisterUserHandlerServer(grpcServer, &userServer)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}

type GRPCServiceHandler struct {
	log logger.Logger
}

func NewGRPCServiceHandler(log logger.Logger) *GRPCServiceHandler {
	return &GRPCServiceHandler{log: log}
}
