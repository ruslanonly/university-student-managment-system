package transport

type Config struct {
	Address                 string        `yaml:"address"`
	GracefulShutdownTimeout int           `yaml:"graceful_shutdown_timeout"`
	ReadTimeout             int           `yaml:"read_timeout"`
	WriteTimeout            int           `yaml:"write_timeout"`
	IdleTimeout             int           `yaml:"idle_timeout"`
	Swagger                 SwaggerConfig `json:"swagger"`
}

type SwaggerConfig struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
	Endpoint string `yaml:"endpoint"`
}
