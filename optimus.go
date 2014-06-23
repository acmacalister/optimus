package main

import (
	"github.com/acmacalister/skittles"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(skittles.Red("please provide environment"))
	}

	environment := os.Args[1]

	loadConfig(environment)
	buildIPA()
	log.Println(skittles.BoldGreen("All done!"))
}
