package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	. "github.com/Financial-Times/golang-app-template/logging"
	"github.com/kr/pretty"
)

// ParseConfig opens the file at configFileName and unmarshals it into an AppConfig.
func ParseConfig(configFileName string) (*AppConfig, error) {
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Printf("Error reading configuration file [%v]: [%v]", configFileName, err.Error())
		return nil, err
	}

	var conf AppConfig
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Printf("Error unmarshalling configuration file [%v]: [%v]", configFileName, err.Error())
		return nil, err
	}

	Info.Printf("Using configuration: %# v", pretty.Formatter(conf))
	return &conf, nil
}
