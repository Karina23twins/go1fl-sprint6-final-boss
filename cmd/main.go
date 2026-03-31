package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logFile, err := os.Create("logs.txt")
	if err != nil {
		log.Fatal("failed to create file:", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "error: ", log.LstdFlags|log.Lshortfile)

	myServ := server.NewServer(logger)

	err = myServ.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}
