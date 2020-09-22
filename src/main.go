package main

import (
	"log"
	"os"

	"github.com/gufranmirza/redact-api-golang/src/models"

	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/web/server"
)

func main() {
	log := log.New(os.Stdout, "main => ", log.LstdFlags)

	// Initialize Application configuration
	_, err := config.LoadConfig(models.DefaultConfigPath)
	if err != nil {
		log.Fatalf("Reading configuration from JSON (%s) failed: %s\n", models.DefaultConfigPath, err)
	}

	// Start http server
	svr := server.NewServer()
	svr.Start()
}
