package main

import (
	"data-hub/nats"
	"data-hub/subscriber"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Print("Initializing application")

	loadEnvError := mapPropertiesToEnv()
	if loadEnvError != nil {
		log.Fatalf("Aborting application start")
		return
	}

	log.Print("Initializing message receiver")
	messageReceiver := &subscriber.Receiver{
		NatsConnection: nats.CreateNatsConnection(),
		MessageSubject: os.Getenv("MESSAGE_SUBJECT"),
	}

	log.Print("Start listening for messages")
	messageReceiver.ReceiveMessageAndProceed()

	// todo

}

func mapPropertiesToEnv() error {
	log.Print("Loading properties to env variables")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file with message: %v", err)
		return err
	}
	log.Print("Loading properties finished")
	return nil
}
