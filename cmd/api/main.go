package main

import (
	"fmt"
	"log"
	"os"
	"packform-backend/src/pkg/config"
	"packform-backend/src/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	if _, ok := os.LookupEnv("APP_ENV"); !ok {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(".env file not found\n")
			os.Exit(2)
		}
	}

	config.LoadEnvironment()
	env := config.Env.Environment
	if env == "" {
		log.Fatal("empty environment\n")
		os.Exit(2)
	}
	fmt.Printf("environment APP_ENV=%s\n", env)

	server := server.New()
	server.Run()
}
