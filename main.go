package main

import (
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open" // Opens file in external editor
	"github.com/tarm/serial"
	"os"
	"time"
)

var port = flag.String("port", "", "What port is the Arduino connected to?")

const (
	FILENAME = "output.txt"
)

func main() {
	flag.Parse()
	if *port == "" {
		fmt.Println("Port is required")
		return
	}
	findArduino()
	export_data("Hello World")
}

func findArduino() {
	config := &serial.Config{Name: *port, Baud: 115200,
		ReadTimeout: time.Second * 5}

	s, err := serial.OpenPort(config)

	if err != nil {
		fmt.Printf("Unable to open serial port: %s", err.Error())
		return
	}

	buf := make([]byte, 10000)
	n, err := s.Read(buf)

	fmt.Printf("Bytes recieved: %d", n)

	if err != nil {
		fmt.Printf("Unable to read serial port: %s", err.Error())
		return
	}
	if n == 0 {
		fmt.Printf("No data recieved")
		return
	}

	fmt.Printf("Data read: \n\n%s\n\n", string(buf[:n]))
}

func export_data(data string) {
	file, err := os.Create(FILENAME)

	if err != nil {
		fmt.Printf("Error opening file: %s", err.Error())
		return
	}

	file.WriteString(data)
	file.Close()
	open.Run(FILENAME)
}
