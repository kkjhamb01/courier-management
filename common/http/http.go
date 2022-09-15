package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func (s Server) ExposeEndpoint() {
	http.Handle("/", s.router)

	port := ":" + config.InternalHttp().Port
	logger.Infof("Internal http server is listening to %v ...", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Error("failed to start http server", err)
		panic(err)
	}
}

func (s Server) AddHealthCheck(readyz func(http.ResponseWriter, *http.Request)) Server {
	s.router.HandleFunc("/livez", livez)
	s.router.HandleFunc("/readyz", readyz)
	return s
}

func (s Server) AddPrometheus(ctx context.Context, setupPrometheus func(ctx context.Context)) Server {
	setupPrometheus(ctx)
	s.router.Handle("/metrics", promhttp.Handler())
	return s
}

func CreateServer() Server {
	return Server{
		router: mux.NewRouter(),
	}
}
