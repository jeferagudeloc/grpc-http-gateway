package infrastructure

import (
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/infrastructure/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	appName    string
	logger     logger.Logger
	ctxTimeout time.Duration
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

func (c *config) Start() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dns:///server:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.logger.Fatalln("could not connect to server: %v", err)
	}

	defer conn.Close()

}
