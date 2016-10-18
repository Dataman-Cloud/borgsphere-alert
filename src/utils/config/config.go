package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr   string                            `yaml:"addr"`
	Input  map[string]map[string]interface{} `yaml:"input"`
	Filter map[string]map[string]interface{} `yaml:"filter"`
	Output map[string]map[string]interface{} `yaml:"output"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		Init()
	}
	return config
}

func Init() {
	envFile := flag.String("config", "env_file.yaml", "")
	if y, err := ioutil.ReadFile(*envFile); err != nil {
		log.Fatalf("read yaml config file error: %v", err)
	} else {
		config = new(Config)
		if err := yaml.Unmarshal(y, config); err != nil {
			log.Fatalf("unmarshal yaml error: %v", err)
		}
	}
}

func ParseConfig(data interface{}, target interface{}) error {
	o, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(o, target)
	if err != nil {
		return err
	}

	return nil
}
