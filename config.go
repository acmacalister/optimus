package main

import (
	"github.com/acmacalister/skittles"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
)

type configuration struct {
	projectName         string
	awsKey              string
	awsSecret           string
	bucket              string
	path                string
	signingIdentity     string
	provisioningProfile string
}

var Config = new(configuration)

func loadConfig(environment string) {
	conf, err := yaml.ReadFile("optimus.yml")

	if err != nil {
		log.Fatal(skittles.Red(err))
	}

	Config.projectName, _ = conf.Get(environment + ".project_name")
	Config.awsKey, _ = conf.Get(environment + ".aws_key")
	Config.awsSecret, _ = conf.Get(environment + ".aws_secret")
	Config.bucket, _ = conf.Get(environment + ".bucket")
	Config.path, _ = conf.Get(environment + ".path")
	Config.signingIdentity, _ = conf.Get(environment + ".signing_identity")
	Config.provisioningProfile, _ = conf.Get(environment + ".provisioning_profile")
}
