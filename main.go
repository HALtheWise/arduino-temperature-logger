package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open" // Opens file in external editor
	"github.com/tarm/serial"
	"os"
	"strings"
	"time"
)

var port = flag.String("port", "", "What port is the Arduino connected to?")

const (
	FILENAME = "output.txt"
	EOM      = "END OF FILE"
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
	config := &serial.Config{Name: *port, Baud: 9600,
		ReadTimeout: time.Second * 5}

	s, err := serial.OpenPort(config)

	if err != nil {
		fmt.Printf("Unable to open serial port: %s", err.Error())
		return
	}

	scanner := bufio.NewScanner(s)

	var lines []string

	var line string
	for line != EOM && len(lines) < 10 {

		success := scanner.Scan()
		line := scanner.Text()

		if !success {
			fmt.Printf("Unable to read scanner: %s", scanner.Err().Error())
			return
		}
		if len(line) == 0 {
			fmt.Printf("No data recieved")
			return
		}

		lines = append(lines, line)

	}

	fmt.Printf("Data read: \n\n%s\n\n", strings.Join(lines, "\n"))
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
