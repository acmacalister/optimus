package main

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
)

type configuration struct {
	projectName string
	awsKey      string
	awsSecret   string
	bucket      string
}

var Config = new(configuration)

func loadConfig(environment string) {
	conf, err := yaml.ReadFile("optimus.yml")

	if err != nil {
		log.Fatal(err)
	}

	Config.projectName, _ = conf.Get(environment + ".project_name")
	Config.awsKey, _ = conf.Get(environment + ".aws_key")
	Config.awsSecret, _ = conf.Get(environment + ".aws_secret")
	Config.bucket, _ = conf.Get(environment + ".bucket")
}
