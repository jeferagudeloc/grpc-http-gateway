module github.com/jeferagudeloc/grpc-http-gateway/src/gateway

go 1.19

require (
	github.com/gorilla/mux v1.8.0
	github.com/jeferagudeloc/grpc-http-gateway/src/server v0.0.0-20230525065623-8e19d6e95f77
	github.com/joho/godotenv v1.5.1
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.9.0
	github.com/sirupsen/logrus v1.9.2
	github.com/urfave/negroni v1.0.0
	google.golang.org/grpc v1.55.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
