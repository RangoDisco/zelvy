package main

import (
	bot "github.com/rangodisco/zelvy/bot/pkg/config/setup"
	"github.com/rangodisco/zelvy/config"
	server "github.com/rangodisco/zelvy/server/config"
	"github.com/rangodisco/zelvy/server/config/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

}

func main() {
	errChan := make(chan error, 2)
	stopChan := make(chan struct{})
	err := database.SetupDatabase()
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}

	server.SetupGRpc(errChan, stopChan)

	bot.Setup(errChan, stopChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case err := <-errChan:
		if err != nil {
			log.Fatalf("failed to start app: %v", err)
		}
	case sig := <-sigChan:
		log.Printf("received signal: %v", sig)
	}

	close(stopChan)
	time.Sleep(2 * time.Second)
}
