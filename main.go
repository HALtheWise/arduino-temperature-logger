package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/skratchdot/open-golang/open" // Opens file in external editor
	"go.bug.st/serial"
)

var FILENAME = flag.String("o", "output.txt", "Output file name")

func main() {
	flag.Parse()
	printHelp()
	port, err := findArduino()
	if err != nil {
		log.Fatal(err)
	}
	readData(port)
	open.Run(*FILENAME)
}

func printHelp() {
	log.Println("This program is designed to record data from an Arduino datalogger to a file.")
	log.Println("Usage:")
	log.Println("\t1. Place the thermocouples in their desired places.")
	log.Println("\t2. Connect the Arduino to the computer.")
	log.Println("\t3. Run this program. Verify that measurements are showing up.")
	log.Println("\t4. When datalogging is complete, unplug the Arduino's USB cable.")
	log.Println("\t5. Copy the resulting text file (default: output.txt) to an Excel workbook.")
	log.Print("\n\n\n")
	time.Sleep(time.Second)
}

func findArduino() (string, error) {
	ports, err := serial.GetPortsList()
	log.Printf("Unsorted list of found ports: %s\n", ports)
	if err != nil {
		return "", err
	}
	if len(ports) == 0 {
		return "", errors.New("No devices found")
	}
	sort.Strings(ports)
	port := ports[len(ports)-1]
	log.Printf("Selected port: %s\n", port)
	return port, nil
}

func readData(port string) {
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	sp, err := serial.OpenPort(port, mode)
	defer sp.Close()

	buff := bufio.NewReader(sp)

	file, _ := os.Create(*FILENAME)
	defer file.Close()
	defer log.Printf("Recieved data saved to %s\n", *FILENAME)

	mwr := io.MultiWriter(file, os.Stdout)

	// Attempt to read the first line, to see if it's broken
	firstline, _ := buff.ReadBytes('\n')
	if !strings.ContainsAny(string(firstline), "03456789") {
		// The first line seems not broken (heuristic)
		log.Printf("First line added back in: \"%s\"\n", strings.TrimSpace(string(firstline)))
		mwr.Write(firstline)

	}

	n, err := io.Copy(mwr, buff)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Bytes read: %d", n)
}
