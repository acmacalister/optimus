package main

import (
	"github.com/acmacalister/skittles"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
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
	out, err = exec.Command("xcodebuild", "-exportArchive", "-exportFormat", "ipa", "-archivePath", archive, "-exportPath", ipa, "-exportProvisioningProfile", Config.provisioningProfileName).CombinedOutput()
	if err != nil {
		log.Fatal(skittles.Red(string(out)))
	}

	uploadIPA(ipa)
}

func uploadIPA(file string) {
	log.Println(skittles.BoldCyan("Uploading ipa..."))
	s := s3.New(aws.Auth{Config.awsKey, Config.awsSecret}, aws.USWest2)
	bucket := s.Bucket(Config.bucket)
	b, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(skittles.Red(err))
	}

	err = bucket.Put(Config.path, b, "text/plain", s3.PublicRead)

	if err != nil {
		log.Fatal(skittles.Red(err))
	}
}
