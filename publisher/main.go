package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1833")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT error: %s", token.Error())
	}

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := client.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	client.Disconnect(250)

	fmt.Println("Complete publish")
}

