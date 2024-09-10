package main

import (
	"bam-catalog/internal/core"
	"log"
)

func main() {
	log.Print("Start application")

	err := core.Run()
	if err != nil {
		log.Fatal(err)
	}
}
