package main

import (
	"os"

	"flag"
	"log"

	. "github.com/Financial-Times/golang-app-template/config"
	. "github.com/Financial-Times/golang-app-template/logging"
)

var configFileName = flag.String("config", "", "Path to configuration file")

func main() {
	InitLogs(os.Stdout, os.Stdout, os.Stderr)
	flag.Parse()

	var err error
	appConfig, err := ParseConfig(*configFileName)
	if err != nil {
		log.Printf("Cannot load configuration: [%v]", err)
		return
	}

	Info.Printf("A value: [%v]", appConfig.Config1)
}
