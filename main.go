package main

import (
	"study_docker/http_server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := http_server.StartHTTPServer(); err != nil {
		panic(err)
	}
}
