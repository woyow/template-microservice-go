package main

import (
	"log"

	"github.com/woyow/example-module/config"
	"github.com/woyow/example-module/internal/app"
)

func main() {
	// Load configuration of application
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// Run application
	app.Run(cfg)
}