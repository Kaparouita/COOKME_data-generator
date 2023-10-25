package main

import (
	"log"
	"os"

	"data-generator/internals/core"
	"data-generator/internals/handlers"
	"data-generator/internals/repositories"

	"github.com/Kaparouita/models/myrabbit/amqp"
	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handler := amqp.AmqpHandlerInstance(os.Getenv("RABBITMQ_DIAL"))

	rc := make(chan bool)
	handler.AutoRedial(os.Getenv("RABBITMQ_DIAL"), rc)

	go func() {
		for ; true; <-rc {
			db := repositories.NewDbRepo(handler)
			srv := core.NewGenerateService(db)
			generateHandler := handlers.NewHandler(srv,handler)
	
			generateHandler.InitServer()
		}
	}()
	// for here to read all plugins
	forever := make(chan bool)

	<-forever
}
