package main

import (
	//"github.com/tarm/serial"
	"fmt"
	"go.bug.st/serial"
	"log"
	//"time"
)

func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
