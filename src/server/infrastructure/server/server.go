package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/grpc/order"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/grpc/user"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/presenter"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/repository"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
)

type grpcServer struct {
	log logger.Logger
	sql repository.SQL
}

func NewGRPCServer(
	log logger.Logger,
	sql repository.SQL,
) *grpcServer {
	return &grpcServer{
		log: log,
		sql: sql,
	}
}

func (s *grpcServer) Listen() {
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	// main server
	grpcServer := grpc.NewServer()

	var (
		uc = usecase.NewCreateOrderInteractor(
			repository.NewMysqlSQL(s.sql),
			presenter.NewCreateOrderPresenter(),
		)
	)

	// server declarations
	orderServer := order.NewServer(uc, nil)
	userServer := user.Server{}

	order.RegisterOrderHandlerServer(grpcServer, orderServer)
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
