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

	port := ports[0]
	recieveComms(port)
}

func recieveComms(port string) {

	mode := &serial.Mode{
		BaudRate: 9600,
	}
	sp, err := serial.OpenPort(port, mode)
	defer sp.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	buff := make([]byte, 100)
	n, err := sp.Read(buff)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buff[:n]))
}
