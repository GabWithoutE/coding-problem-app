package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *zap.Logger
}

type Server interface {
	SetupRoutes()
	Logger() *zap.Logger
	Run() error
}

func NewServer() (Server, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, errors.Wrap(err, "NewServer: failed to initialize logger")
	}

	return &server{
		router: mux.NewRouter(),
		logger: logger,
	}, nil
}

func (s *server) Logger() *zap.Logger {
	return s.logger
}

func (s *server) Run() error {
	server := &http.Server{Addr: ":8080", Handler: s.router}

	s.logger.Info("server started", zap.String("listner", "service"), zap.String("addr", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		serErr := server.Shutdown(context.Background())
		return multierr.Combine(
			err,
			serErr,
		)
	}

	return nil
}
