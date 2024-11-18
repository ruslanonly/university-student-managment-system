package logger

type Config struct {
	Level string `yaml:"level" default:"debug"` // Уровень логирования
}
