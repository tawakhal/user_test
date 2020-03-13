package main

import (
	slog "github.com/tawakhal/user_test/extend/logger"
	shv "github.com/tawakhal/user_test/sharevar"
)

func main() {
	// Create Logging
	shv.Logger = slog.NewLogger(slog.Option{
		LogLevel: slog.LevelALL,
		Format:   slog.FormatFMT,
		Ouput:    slog.Console,
	})
	shv.Logger.Info("Server Started")
}
