package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gufranmirza/redact-api-golang/src/models"
)

// Config is singleton config object
var Config *models.AppConfig

// LoadConfig reads and loads configuration if it was not read already
func LoadConfig(configFilePath string) (*models.AppConfig, error) {
	var err error
	if Config == nil {
		Config, err = ReadConfigFromJSON(configFilePath)
	}
	return Config, err
}

// ReadConfigFromJSON reads config data from JSON-file
func ReadConfigFromJSON(path string) (*models.AppConfig, error) {
	log := log.New(os.Stdout, "config :=> ", log.LstdFlags)
	log.Printf("Looking for JSON config file (%s)", path)

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", path, err)
		return nil, err
	}

	reader := bytes.NewBuffer(contents)
	err = json.NewDecoder(reader).Decode(&Config)
	if err != nil {
		return nil, err
	}

	log.Printf("Configuration has been read from JSON (%s) successfully\n", path)
	return Config, nil
}
