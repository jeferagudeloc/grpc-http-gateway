package main

import (
	"os"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/database"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/log"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/infrastructure/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		SqlSetup(database.InstanceMysql)

	app.GrpcServer(server.InstanceGRPC).StartGrpc()

}
