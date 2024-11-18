package transport

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	authHandler "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/handler"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTPServer struct {
	router      *chi.Mux
	log         *slog.Logger
	cfg         *Config
	authHandler *authHandler.Handler
}

func (s *HTTPServer) useSwagger() {
	s.router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("Swagger Protected Area", map[string]string{
			s.cfg.Swagger.Login: s.cfg.Swagger.Password,
		}))

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", s.cfg.Swagger.Endpoint)),
		))
	})

	s.log.Info(fmt.Sprintf("Visit %s/swagger/index.html for documentation.", s.cfg.Swagger.Endpoint))
}

func (s *HTTPServer) initRoutes() {
	authHandler.Init(s.router, s.authHandler)
}

func (s *HTTPServer) serve() {
	server := &http.Server{
		Addr:         s.cfg.Address,
		Handler:      s.router,
		WriteTimeout: time.Duration(s.cfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(s.cfg.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.cfg.IdleTimeout) * time.Second,
	}

	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-done

		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(s.cfg.GracefulShutdownTimeout)*time.Second,
		)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			s.log.Error("HTTP shutdown error: %v", err)
			return
		}
	}()

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		s.log.Error("HTTP server error: %v", err)
		return
	}
}

func (s *HTTPServer) Run() {
	s.initRoutes()
	s.useSwagger()
	s.serve()
}

func NewHTTPServer(log *slog.Logger, cfg *Config, authHandler *authHandler.Handler) *HTTPServer {
	return &HTTPServer{
		router:      chi.NewRouter(),
		log:         log,
		cfg:         cfg,
		authHandler: authHandler,
	}
}
