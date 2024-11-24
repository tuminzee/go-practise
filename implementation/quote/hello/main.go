package main

import (
	"fmt"
	"log"

	greet "go/gretting"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	msg, err := greet.Hello("Yolo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

	names := []string{"Tumin", "Yolo", "Tunde"}
	msgs, err := greet.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msgs)

	// msg, err = greet.Hello("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(msg)
}
