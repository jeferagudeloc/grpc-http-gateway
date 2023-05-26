package main

import (
	"os"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/database"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/grpc"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/http"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/log"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		SqlSetup(database.InstanceMysql).
		GrpcServer(grpc.InstanceGRPC).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(http.InstanceGorillaMux)
	app.StartServers()
}
