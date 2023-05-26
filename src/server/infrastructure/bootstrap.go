package infrastructure

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/repository"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/database"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/grpc"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/http"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/log"
)

type config struct {
	appName       string
	logger        logger.Logger
	ctxTimeout    time.Duration
	dbSQL         repository.SQL
	grpcServer    grpc.Server
	webServer     http.Server
	webServerPort http.Port
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

func (c *config) SqlSetup(instance int) *config {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		c.logger.Fatalln(err, "There was an error setting the database")
	}

	c.logger.Infof("Successfully connected to the SQL database")

	c.dbSQL = db
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

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.webServerPort = http.Port(p)
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := http.NewWebServerFactory(
		instance,
		c.logger,
		c.webServerPort,
		c.ctxTimeout,
		c.dbSQL,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) GrpcServer(instance int) *config {
	s, err := grpc.NewGrpcServerFactory(
		instance,
		c.logger,
		c.dbSQL,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured grpc server")

	c.grpcServer = s
	return c
}

func (c *config) StartServers() {
	go func() {
		c.grpcServer.Listen()
	}()

	go func() {
		c.webServer.Listen()
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

}
