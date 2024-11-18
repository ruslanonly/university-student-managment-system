package sl

import (
	"github.com/ruslanonly/university-student-managment-system/src/pkg/logger"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

// Функция для Pretty лога (для консоли в режиме разработки)
func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

// Err возвращает атрибут ошибки
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Op(name string) slog.Attr {
	return slog.Attr{
		Key:   "operation",
		Value: slog.StringValue(name),
	}
}

// SetupLogger настраивает logger
func SetupLogger(cfg logger.Config) *slog.Logger {
	return setupPrettySlog()
}
