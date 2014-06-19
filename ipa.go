package main

import (
	"fmt"
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

	out, err := exec.Command("xcodebuild", "-workspace", workspace, "-scheme", strings.Title(Config.projectName), "archive", "-archivePath", archive).CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
	fmt.Println(string(out))

	out, err = exec.Command("xcodebuild", "-exportArchive", "-exportFormat", "ipa", "-archivePath", archive, "-exportPath", ipa, "-exportProvisioningProfile", "My Big Campus App Beta").CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
	fmt.Println(string(out))

	uploadIPA(ipa)
}

func uploadIPA(file string) {
	s := s3.New(aws.Auth{Config.awsKey, Config.awsSecret}, aws.USWest2)
	bucket := s.Bucket(Config.bucket)
	b, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	err = bucket.Put("path", b, "text/plain", s3.PublicRead)

	if err != nil {
		log.Fatal(err)
	}
}
