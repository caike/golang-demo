package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Hello")
		return
	}

	message := os.Args[1]

	f, errCreate := os.Create("odevs.txt")
	if errCreate != nil {
		log.Fatal("Error opening file") // returns exit code 1
	}
	defer f.Close()

	_, errWrite := f.WriteString(message)
	if errWrite != nil {
		log.Fatal("Error opening file") // returns exit code 1
	}
	f.Sync()
	fmt.Println("Message written to odevs.txt")
}
