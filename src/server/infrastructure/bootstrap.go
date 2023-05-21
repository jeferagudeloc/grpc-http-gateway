package infrastructure

import (
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/log"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/server"
)

type config struct {
	appName    string
	logger     logger.Logger
	ctxTimeout time.Duration
	server     server.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) Name(name string) *config {
	c.appName = name
	return c
}

func (c *config) Logger(instance int) *config {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Successfully configured log")
	return c
}

func (c *config) GrpcServer(instance int) *config {
	s, err := server.NewGrpcServerFactory(
		instance,
		c.logger,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured grpc server")

	c.server = s
	return c
}

func (c *config) StartGrpc() {
	c.server.Listen()
}
