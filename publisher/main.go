package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT error: %s", token.Error())
	}

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := client.Publish("net/sample", 0, false, text)
		token.Wait()
	}

	client.Disconnect(250)

	fmt.Println("Complete publish")
}
