package config

import (
	"log/slog"
	"os"
)

func setLogger() {
	hanlder := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug | slog.LevelError | slog.LevelInfo | slog.LevelWarn,
	})
	slog.SetDefault(slog.New(hanlder))
}

func init() {
	setLogger()
}
