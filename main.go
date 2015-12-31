package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/skratchdot/open-golang/open" // Opens file in external editor
	"go.bug.st/serial"
)

var FILENAME = flag.String("o", "output.txt", "Output file name")

func main() {
	flag.Parse()
	port, err := findArduino()
	if err != nil {
		log.Fatal(err)
	}
	readData(port)
	open.Run(*FILENAME)
}

func findArduino() (string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return "", err
	}
	if len(ports) == 0 {
		return "", errors.New("No devices found")
	}
	sort.Strings(ports)
	return ports[len(ports)-1], nil
}

func readData(port string) {
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	sp, err := serial.OpenPort(port, mode)
	defer sp.Close()

	buff := bufio.NewReader(sp)
	buff.ReadLine()

	file, _ := os.Create(*FILENAME)
	defer file.Close()

	mwr := io.MultiWriter(file, os.Stdout)

	n, err := io.Copy(mwr, buff)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bytes read: %d", n)
}
