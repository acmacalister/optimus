package main

import (
	"fmt"
	"log"
	"os"
)

const (
	Archive   = ".xcarchive"
	IPA       = ".ipa"
	Workspace = ".xcworkspace"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("please provide name of xcode project and aws bucket name")
	}

	name := os.Args[1]
	bucketName := os.Args[2]
	fmt.Println(name, bucketName)

	buildIPA(name)
}
