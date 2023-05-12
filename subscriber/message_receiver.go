package subscriber

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type Receiver struct {
	NatsConnection *nats.Conn
	MessageSubject string
}

func (receiver *Receiver) ReceiveMessageAndProceed() {

	_, subscribeError := receiver.NatsConnection.Subscribe(receiver.MessageSubject, func(message *nats.Msg) {
		// need to implement further
		receiver.proceedMessage(message)
	})
	if subscribeError != nil {
		log.Fatal(subscribeError)
	}

	// Keep the application running indefinitely
	select {}
}

func (receiver *Receiver) proceedMessage(message *nats.Msg) {
	fmt.Printf("Message: %v", message)
}
