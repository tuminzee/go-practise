package main

import (
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_publisher")

	// client := MQTT.NewClient(opts)
	// if token := client.Connect(); token.Wait() && token.Error() != nil {
	// 	log.Fatal(token.Error())
	// }

	// defer client.Disconnect(250)

	// for {
	// 	message := fmt.Sprintf("Message at %v", time.Now())
	// 	token := client.Publish("test/topic", 0, false, message)
	// 	token.Wait()
	// 	fmt.Printf("Published: %s\n", message)
	// 	time.Sleep(5 * time.Second)
	// }

	// Create a new MQTT client
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	defer client.Disconnect(250)

	// Message payload
	message := "Hello, MQTT!"

	// Publish one million messages
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		payload := fmt.Sprintf("%s %d", message, i)
		token := client.Publish("test/topic", 0, false, payload)
		token.Wait() // Wait for the message to be published
		if token.Error() != nil {
			log.Printf("Error publishing message %d: %v", i, token.Error())
		}
		if i%10000 == 0 {
			fmt.Printf("Published %d messages\n", i)
		}
	}
	duration := time.Since(start)
	fmt.Printf("Published 1,000,000 messages in %v\n", duration)
}
