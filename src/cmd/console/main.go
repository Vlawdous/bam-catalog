package main

import (
	"bam-catalog/internal/core"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Print("Запуск команды: " + strings.Join(os.Args, " "))

	err := core.RunCli()
	if err != nil {
		log.Fatal("Ошибка выполнения команды: " + err.Error())
	}
}
