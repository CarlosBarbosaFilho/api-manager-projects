package main

import (
	"github.com/CarlosBarbosaFilho/api-manager-projects/config"
	"github.com/CarlosBarbosaFilho/api-manager-projects/router"
)

var (
	logger *config.Logger
)

func main() {

	// Initializer logger
	logger = config.GetLogger(">>>>> main <<<<<")

	// Initialize configs
	config.Init()

	// Initialize routes
	router.Initialize()
}
