package main

import (
	"fmt"

	"github.com/CarlosBarbosaFilho/api-manager-projects/config"
	"github.com/CarlosBarbosaFilho/api-manager-projects/router"
)

var (
	logger *config.Logger
)

func main() {
	fmt.Println("Initializer Application !")

	// Initializer logger
	logger = config.GetLogger(">>>>> main <<<<<")

	// Initialize configs
	config.Init()

	// Initialize routes
	router.Initialize()
}
