package app

import (
	"context"
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	_ "github.com/ruslanonly/university-student-managment-system/src/docs"
	"github.com/ruslanonly/university-student-managment-system/src/internal/transport"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/logger/sl"
	"log"
	"log/slog"
	"os"
)

// App
// @title University Report Generation System
// @version 0.1
// @BasePath /
type App struct {
	cfg        *config
	log        *slog.Logger
	router     *chi.Mux
	container  *diContainer
	httpServer *transport.HTTPServer
}

func (a *App) mustLoadConfig() {

	configPath := flag.String("config", "", "path to yaml configure file.")

	flag.Parse()

	if *configPath == "" {
		log.Fatalf("config command flag is not set")
	}

	if _, err := os.Stat(*configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	cfg := new(config)

	if err := cleanenv.ReadConfig(*configPath, cfg); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	a.cfg = cfg
}

func (a *App) setupLogger() {
	a.log = sl.SetupLogger(a.cfg.Logger)
}

func (a *App) useCookies() {
	cfg := a.cfg.Cookie
	store := sessions.NewCookieStore([]byte(cfg.Secret))
	store.MaxAge(cfg.MaxAge)
	gothic.Store = store
}

func (a *App) useProviders() {
	cfg := a.cfg.Providers

	goth.UseProviders(
		google.New(
			cfg.Google.ClientKey,
			cfg.Google.Secret,
			cfg.Google.CallbackURL,
			cfg.Google.Scopes...,
		),
	)
}

func (a *App) serveHTTP() {
	httpServer := transport.NewHTTPServer(a.log, &a.cfg.HTTPServer,
		a.container.authService,
		a.container.authHandler,
		a.container.lab1Handler,
		a.container.lab2Handler,
		a.container.lab3Handler,
	)
	a.httpServer = httpServer
	httpServer.Run()
}

func (a *App) Run() {
	a.mustLoadConfig()
	a.setupLogger()

	a.log.Info("Config", a.cfg)

	a.useCookies()
	a.useProviders()
	cleanup := a.inject(context.Background())
	defer cleanup()
	a.serveHTTP()
}

func New() *App {
	return &App{}
}
