package main

import (
	"fmt"

	"github.com/ElielClementino/Go_end/config"
	"github.com/ElielClementino/Go_end/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errf("Config initialization error: %v", err)
		fmt.Println(err)
		return
	}

	// Initialize the router listening on :8080
	router.InitializeRouter()
}
