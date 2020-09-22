package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gufranmirza/redact-api-golang/src/web/router"

	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/models"
	"github.com/gufranmirza/redact-api-golang/src/web/middlewares"
)

type server struct {
	logger            *log.Logger
	startTimestampUTC time.Time
	config            *models.AppConfig
	router            router.Router
}

// NewServer returns a new instance of a server
func NewServer() Server {
	return &server{
		logger: log.New(os.Stdout, "server :=> ", log.LstdFlags),
		config: config.Config,
		router: router.NewRouter(),
	}
}

// Start starts a new HTTP server and register routes
func (server *server) Start() {
	server.logger.Println("Server is starting up...")
	addr := fmt.Sprintf("%s:%v", server.config.Hostname, server.config.Port)

	// initialize router
	r := server.router.Router()

	// initialize server
	s := &http.Server{
		Addr:         addr,
		Handler:      middlewares.Tracing(server.logger)(middlewares.Logging(server.logger)(r)),
		ErrorLog:     server.logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		server.logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.SetKeepAlivesEnabled(false)
		if err := s.Shutdown(ctx); err != nil {
			server.logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	server.logger.Println("Server is ready to handle requests at: ", addr)
	server.startTimestampUTC = time.Now().UTC()
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		server.logger.Fatalf("Could not listen on %s: %v\n", addr, err)
	}

	<-done
	server.logger.Println("Server stopped")
}

func (server *server) StartTimestampUTC() time.Time {
	return server.startTimestampUTC
}
