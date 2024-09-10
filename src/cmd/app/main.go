package main

import (
	"bam-catalog/internal/core"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Print("Запуск приложения")

	err := core.RunHttp()
	if err != nil {
		log.Fatal("Не удалось запустить приложение: " + err.Error())
	}
}
