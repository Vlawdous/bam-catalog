package main

import (
	"bam-catalog/internal/app"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Start command")

	err := app.StartCommand()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Command executed successfully")
}
