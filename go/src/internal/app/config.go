package app

import (
	"github.com/ruslanonly/university-student-managment-system/src/internal/transport"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/logger"
)

type providerConfig struct {
	ClientKey   string   `yaml:"client_key"`
	Secret      string   `yaml:"secret"`
	CallbackURL string   `yaml:"callback_url"`
	Scopes      []string `yaml:"scopes"`
}

type providersConfig struct {
	Google providerConfig `yaml:"google"`
}

type cookieConfig struct {
	Secret string `yaml:"secret"`
	MaxAge int    `yaml:"max_age"`
}

type config struct {
	HTTPServer transport.Config `yaml:"http_server"`
	Logger     logger.Config    `yaml:"logger"`
	Providers  providersConfig  `yaml:"providers"`
	Cookie     cookieConfig     `yaml:"cookie"`
}
