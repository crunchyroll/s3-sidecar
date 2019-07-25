package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// AppConfig - holds configuration details for the App
type AppConfig struct {
	Logging  LoggingConfig  `yaml:"logging"`
	Newrelic NewrelicConfig `yaml:"newrelic"`
	S3       S3Config       `yaml:"s3"`
	Service  ServiceConfig  `yaml:"service"`
}

// LoggingConfig - holds configuration details for the logging lib
type LoggingConfig struct {
	AppName    string `yaml:"app_name"`
	AppVersion string `yaml:"app_version"`
	EngGroup   string `yaml:"eng_group"`
	Level      string `yaml:"level"`
}

// ServiceConfig - contains all the service details for service
type ServiceConfig struct {
	Listen int `yaml:"listen"`
}

// NewrelicConfig - contains all key details for invoking Newrelic go sdk
type NewrelicConfig struct {
	License string `yaml:"license"`
	Name    string `yaml:"name"`
}

// S3Config - contains all key details for invoking Newrelic go sdk
type S3Config struct {
	Region string `yaml:"region"`
	Bucket string `yaml:"bucket"`
}

// LoadConfiguration - read configurations from a yaml file and loads into 'Config' struct.
//					   returns error if the file is missing or contains bad schema.
func LoadConfiguration(filePath string) (*AppConfig, error) {
	cnf := &AppConfig{}

	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read configuration file %s", err.Error())
	}

	ymlErr := yaml.Unmarshal(rawContent, &cnf)
	if ymlErr != nil {
		return nil, fmt.Errorf("unable to unpack file %s", ymlErr.Error())
	}

	return cnf, nil
}
