package main

import (
	"api/routes"
	"errors"
	"log"
	"os"
)

const (
	pathFiles = "resources"
)

func main() {
	//make directory if not exists
	if _, err := os.Stat(pathFiles); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFiles, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	routes.Setup()
}
