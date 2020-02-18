package main

import (
	shv "github.com/tawakhal/user_test/sharevar"
	"github.com/tawakhal/util/logging"
)

func main() {
	// Create Logging
	shv.Logger = logging.NewLogger(logging.Option{
		LogLevel: logging.LevelALL,
		Format:   logging.FormatFMT,
		Ouput:    logging.Console,
	})
	shv.Logger.Info("Server Started")
}
