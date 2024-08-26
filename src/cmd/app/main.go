package main

import (
	"bam-catalog/internal/app"
	"log"
)

func main() {
	log.Print("Start application")

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
