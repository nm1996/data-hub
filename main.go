package main

import (
	"data-hub/config"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	natsConfiguration := config.GetNatsConfiguration()

	// Connect to NATS
	nc, err := nats.Connect(natsConfiguration.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to a subject and specify a callback function
	_, err = nc.Subscribe("adv_msg", func(msg *nats.Msg) {
		// Log the message data to the console
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Keep the application running indefinitely
	select {}
}
