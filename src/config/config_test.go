package config

import (
	"fmt"
	"testing"

	"github.com/gufranmirza/redact-api-golang/src/models"
)

var configFilePath = fmt.Sprintf(".%s", models.DefaultConfigPath)

func TestReadConfigFromJSON(t *testing.T) {
	value, err := ReadConfigFromJSON(configFilePath)

	if err != nil {
		t.Errorf("Reading configuration failed: %s\n", err)
	}

	if value == nil {
		t.Errorf("Configuration file is empty, value: %v\n", value)
	}
}

func TestLoadConfig(t *testing.T) {
	Config = nil
	value1, err := LoadConfig(configFilePath)
	value2, err1 := LoadConfig(configFilePath)

	if err != nil {
		t.Errorf("Load configuration failed: %s\n", err)
	}
	if err1 != nil {
		t.Errorf("Load configuration for 2-nd time failed: %s\n", err)
	}

	if value1 != value2 {
		t.Errorf("Configuration file changed")
	}

}
