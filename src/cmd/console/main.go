package main

import (
	"bam-catalog/internal/core"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Start command")

	err := core.StartCommand()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Command executed successfully")
}
