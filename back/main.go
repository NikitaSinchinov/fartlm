package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"main/api_client"
	"main/api_server"

	"go-shared/logger"
	"go-shared/utils"
)

func main() {
	appPort := os.Getenv("BACK_PORT")
	if appPort == "" {
		logger.Fatal("BACK_PORT environment variable is required")
	}

	externalAPIPort := os.Getenv("BACK_EXTERNAL_API_PORT")
	if appPort == "" {
		logger.Fatal("BACK_EXTERNAL_API_PORT environment variable is required")
	}

	app := fiber.New()

	// Middleware
	app.Use(fiber_logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Serve static files
	app.Static("/", "./public")

	// Setup fart client
	hc := http.Client{}
	client, err := api_client.NewClientWithResponses(
		fmt.Sprintf("http://host.docker.internal:%s", externalAPIPort),
		api_client.WithHTTPClient(&hc),
	)
	if err != nil {
		logger.Fatal("%v", err)
	}

	// Setup server
	audioProcessor := NewFartMLAudioProcessor(client)
	server := NewServer(audioProcessor)
	handler := api_server.NewStrictHandler(server, nil)
	api_server.RegisterHandlers(app, handler)

	utils.StartAndListenToSigTerm(
		func() {
			addr := fmt.Sprintf(":%s", appPort)
			logger.Info("Starting server on %s", addr)
			if err := app.Listen(addr); err != nil && err != http.ErrServerClosed {
				logger.Fatal("Listen error: %v", err)
			}
		},
		func() {
			logger.Info("Shutting down server")

			server.Stop()

			if err := app.ShutdownWithTimeout(15 * time.Second); err != nil {
				logger.Fatal("Server forced to shutdown: %v", err)
			}
			logger.Info("Server shutdown complete")
		},
	)
}
