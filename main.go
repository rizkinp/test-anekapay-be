package main

import (
	"log"
	server "test-anekapay-backend/cmd"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
