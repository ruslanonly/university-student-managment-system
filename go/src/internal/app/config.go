package app

import (
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth"
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

type Postgres struct {
	ConnectionString string `yaml:"connection_string"`
}

type Elastic struct {
	ConnectionString string `yaml:"connection_string"`
}

type Neo struct {
	URI      string `yaml:"uri"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Redis struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Mongo struct {
	Address string `yaml:"address"`
}

type config struct {
	HTTPServer transport.Config `yaml:"http_server"`
	Logger     logger.Config    `yaml:"logger"`
	Providers  providersConfig  `yaml:"providers"`
	Cookie     cookieConfig     `yaml:"cookie"`
	Auth       auth.Config      `yaml:"auth"`
	Postgres   Postgres         `yaml:"postgres"`
	Elastic    Elastic          `yaml:"elastic"`
	Neo        Neo              `yaml:"neo"`
	Redis      Redis            `yaml:"redis"`
	Mongo      Mongo            `yaml:"mongo"`
}
