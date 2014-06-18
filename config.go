package main

type configuration struct {
	Enviroment string
	AwsKey     string
	AwsSecret  string
	bucket     string
}

var Config = new(configuration)

func loadConfig() {
	//nothing yet
}
