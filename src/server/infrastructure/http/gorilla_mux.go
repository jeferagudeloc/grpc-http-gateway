package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/api/action"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/api/middleware"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/repository"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
	"github.com/urfave/negroni"
)

type gorillaMux struct {
	router     *mux.Router
	middleware *negroni.Negroni
	log        logger.Logger
	port       Port
	sqlDB      repository.SQL
	ctxTimeout time.Duration
}

func newGorillaMux(
	log logger.Logger,
	port Port,
	t time.Duration,
	sqlDB repository.SQL,
) *gorillaMux {
	return &gorillaMux{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		log:        log,
		port:       port,
		ctxTimeout: t,
		sqlDB:      sqlDB,
	}
}

func (g gorillaMux) Listen() {
	g.setAppHandlers(g.router)
	g.middleware.UseHandler(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.middleware,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server On %s", g.port)
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
	api.Handle("/orders", g.buildFindOrdersAction()).Methods(http.MethodGet)
	api.Handle("/users", g.buildFindUsersAction()).Methods(http.MethodGet)
}

func (g gorillaMux) buildFindOrdersAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewFindOrdersInteractor(
				repository.NewMysqlSQL(g.sqlDB),
			)
			act = action.NewFindOrdersAction(uc, g.log)
		)

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

func (g gorillaMux) buildFindUsersAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewFindUsersInteractor(
				repository.NewMysqlSQL(g.sqlDB),
			)
			act = action.NewFindUsersAction(uc, g.log)
		)

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}
