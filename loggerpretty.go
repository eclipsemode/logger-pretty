package logger_pretty

import (
	"github.com/eclipsemode/logger-pretty/internal/app/slogpretty"
	"log/slog"
	"os"
)

func NewPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
