package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide environment")
	}

	environment := os.Args[1]

	loadConfig(environment)
	buildIPA()
}
