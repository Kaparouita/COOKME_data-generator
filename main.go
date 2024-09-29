package main

import (
	"data-generator/internals/core"
	"data-generator/internals/repositories"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := repositories.NewDbRepo()
	srv := core.NewGenerateService(db)
	// apiKey1 := os.Getenv("PEXELS_API_KEY1")
	apiKey2 := os.Getenv("PEXELS_API_KEY2")
	apiKey3 := os.Getenv("PEXELS_API_KEY3")
	apiKeys := []string{apiKey2, apiKey3}
	for _, apiKey := range apiKeys {
		err := srv.AddImages("", apiKey)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// func main() {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	rc := make(chan bool)

// 	go func() {
// 		for ; true; <-rc {
// 			db := repositories.NewDbRepo(handler)
// 			srv := core.NewGenerateService(db)
// 			generateHandler := handlers.NewHandler(srv, handler)

// 			generateHandler.InitServer()
// 		}
// 	}()
// 	// for here to read all plugins
// 	forever := make(chan bool)

// 	<-forever
// }
