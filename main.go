package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open" // Opens file in external editor
	"os"
)

const (
	FILENAME = "output.tsv"
)

func main() {
	fmt.Println("Hello World")
	export_data("Hello World")
}

func export_data(data string) {
	file, err := os.Create(FILENAME)

	fmt.Printf("File: %s", file.Name())
	if err != nil {
		fmt.Errorf("Error opening file: %s", err.Error())
	}

	file.WriteString(data)
	file.Close()
	open.Run(FILENAME)
}
