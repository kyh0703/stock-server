package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/controller"
	_ "github.com/kyh0703/stock-server/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server client server

// @contact.name API support
// @contact.url http://www.swagger.io/support
// @contact.email kyh0703@nate.com

// @host localhost:8000
// @BasePath /api/v1

func main() {
	// Create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// init ent client
	ec, err := config.ConnectDatabase(ctx)
	if err != nil {
		log.Fatalf("failed connection database: %v", err)
	}
	defer ec.Close()

	// init redis client
	rc, err := config.ConnectRedis()
	if err != nil {
		log.Fatalf("failed connection redis: %v", err)
	}
	defer rc.Close()

	// set routing
	router := controller.NewRouter(ec, rc)
	controller.SetupRouter(router)
	// server configure
	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	// initializing the server in goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal notify user of shutdown.
	stop()

	// The context is used to inform the server it has 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server Exit")
}
