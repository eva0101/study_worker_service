package main

import "study_docker/http_server"

func main() {
	if err := http_server.StartHTTPServer(); err != nil {
		panic(err)
	}
}
