package main

import (
	"os"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/infrastructure"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/infrastructure/log"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger)

	app.Start()

}
