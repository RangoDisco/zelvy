package main

import (
	"github.com/rangodisco/zelvy/server/config"
	"github.com/rangodisco/zelvy/server/config/database"
	"log"
)

func init() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

}

func main() {
	err := database.SetupDatabase()
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}

	err = config.SetupGRpc()
	if err != nil {
		log.Fatalf("failed to setup gRpc: %v", err)
	}
	//err = bot.Setup()
	//if err != nil {
	//	log.Fatalf("failed to setup bot: %v", err)
	//}

}
