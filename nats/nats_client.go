package nats

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func getConnectionProperty() string {
	return os.Getenv("NATS_PORT")
}

func CreateNatsConnection() *nats.Conn {
	log.Printf("Trying to establish connection to nats")
	connection, natsConnectionError := nats.Connect(getConnectionProperty())

	if natsConnectionError != nil {
		log.Fatalf("Connection to nats not established: %v", natsConnectionError)
		panic("Application can't work without nats")
	}

	return connection
}
