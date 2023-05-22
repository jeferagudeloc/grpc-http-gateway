package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/api/action"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/api/middleware"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/presenter"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/usecase"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/urfave/negroni"
)

type gorillaMux struct {
	router     *mux.Router
	middleware *negroni.Negroni
	log        logger.Logger
	ctxTimeout time.Duration
}

func newGorillaMux(
	log logger.Logger,
	t time.Duration,
) *gorillaMux {
	return &gorillaMux{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		log:        log,
		ctxTimeout: t,
	}
}

func (g gorillaMux) Listen() {
	g.setAppHandlers(g.router)
	g.middleware.UseHandler(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      g.middleware,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		g.log.WithFields(logger.Fields{"port": 8080}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			g.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		g.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	g.log.Infof("Service down")
}

func (g gorillaMux) setAppHandlers(router *mux.Router) {
	api := router
	api.HandleFunc("/health", action.HealthCheck).Methods(http.MethodGet)
	api.Handle("/orders", g.buildCreateOrderAction()).Methods(http.MethodPost)
}

func (g gorillaMux) buildCreateOrderAction() *negroni.Negroni {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dns:///server:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		g.log.Fatalln("could not connect to server: %v", err)
	}

	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewCreateOrderInteractor(
				presenter.NewCreateOrderPresenter(),
				g.ctxTimeout,
				g.log,
				conn,
			)
			act = action.NewCreateOrderAction(uc, g.log)
		)

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}
