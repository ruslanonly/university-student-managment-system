package transport

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	authService "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/service"
	authHandler "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/handler"
	lab1Handler "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab1/handler"
	lab2Handler "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab2/handler"
	lab3Handler "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab3/handler"
)

type HTTPServer struct {
	router      *chi.Mux
	log         *slog.Logger
	cfg         *Config
	authService *authService.AuthService
	authHandler *authHandler.Handler
	lab1Handler *lab1Handler.Handler
	lab2Handler *lab2Handler.Handler
	lab3Handler *lab3Handler.Handler
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
	authMiddleware := authHandler.CreateAuthMiddleware(s.authService)
	authHandler.Init(s.router, s.authHandler)
	lab1Handler.Init(s.router, s.lab1Handler, authMiddleware)
	lab2Handler.Init(s.router, s.lab2Handler, authMiddleware)
	lab3Handler.Init(s.router, s.lab3Handler, authMiddleware)
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

func NewHTTPServer(log *slog.Logger, cfg *Config, authService *authService.AuthService, authHandler *authHandler.Handler, lab1Handler *lab1Handler.Handler, lab2Handler *lab2Handler.Handler, lab3Handler *lab3Handler.Handler) *HTTPServer {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)
	return &HTTPServer{
		router:      r,
		log:         log,
		cfg:         cfg,
		authService: authService,
		authHandler: authHandler,
		lab1Handler: lab1Handler,
		lab2Handler: lab2Handler,
		lab3Handler: lab3Handler,
	}
}
