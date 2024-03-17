package main

import (
	"data-generator/internals/core"
	"fmt"
)

func main() {

	srv := core.NewGenerateService(nil)
	err := srv.AddImages("./recipes.json")
	if err != nil {
		fmt.Println(err)
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
