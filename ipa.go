package main

import (
	"github.com/acmacalister/skittles"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	Archive   = ".xcarchive"
	IPA       = ".ipa"
	Workspace = ".xcworkspace"
)

func buildIPA() {
	archive := Config.projectName + Archive
	ipa := Config.projectName + IPA
	workspace := Config.projectName + Workspace

	log.Println(skittles.BoldCyan("Building archive..."))
	out, err := exec.Command("xcodebuild", "-workspace", workspace, "-scheme", strings.Title(Config.projectName), "archive", "-archivePath", archive).CombinedOutput()
	if err != nil {
		log.Fatal(skittles.Red(string(out)))
	}

	log.Println(skittles.BoldCyan("Building ipa..."))
	out, err = exec.Command("xcodebuild", "-exportArchive", "-exportFormat", "ipa", "-archivePath", archive, "-exportPath", ipa, "-exportSigningIdentity", Config.signingIdentity).CombinedOutput()
	if err != nil {
		log.Fatal(skittles.Red(string(out)))
	}

	uploadIPA(ipa)
}

func uploadIPA(file string) {
	log.Println(skittles.BoldCyan("Uploading ipa..."))
	s := s3.New(aws.Auth{Config.awsKey, Config.awsSecret}, aws.USEast)
	bucket := s.Bucket(Config.bucket)
	b, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(skittles.Red(err))
	}

	err = bucket.Put("/"+Config.path+"/"+Config.projectName+IPA, b, "application/x-itunes-ipa", s3.PublicRead)
	if err != nil {
		log.Fatal(skittles.Red(err))
	}

	err = os.RemoveAll(Config.projectName + Archive)
	if err != nil {
		log.Fatal(skittles.Red(err))
	}

	err = os.Remove(Config.projectName + IPA)
	if err != nil {
		log.Fatal(skittles.Red(err))
	}
}
