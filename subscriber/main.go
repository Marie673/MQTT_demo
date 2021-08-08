package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	msgCh := make(chan mqtt.Message)
	var handle mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
		msgCh <- message
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT error: %s", token.Error())
	}

	if subscribeToken := client.Subscribe("net/sample", 0, handle); subscribeToken.Wait() && subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	for {
		select {
		case m := <-msgCh:
			fmt.Printf("topic: %v, payload: %v\n", m.Topic(), string(m.Payload()))
		case <-signalCh:
			fmt.Printf("Interrupt detected.\n")
			client.Disconnect(1000)
			return
		}
	}
}
