package main

import (
	shv "github.com/tawakhal/user_test/sharevar"
	"github.com/tawakhal/util/log"
)

func main() {
	// Create Logging
	shv.Logger = log.NewLogger()
}
