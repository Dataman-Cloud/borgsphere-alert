package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/ghodss/yaml"
)

type Config struct {
	Addr   string                              `yaml:"addr"`
	Input  map[string]map[string]interface{}   `yaml:"input"`
	Filter []map[string]map[string]interface{} `yaml:"filter"`
	Output map[string]map[string]interface{}   `yaml:"output"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		Init()
	}
	return config
}

func Init() {
	envFile := flag.String("config", "config.yaml", "")
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

func (c *Config) GetFilterByModule(module string) (map[string]interface{}, bool) {
	for _, v := range c.Filter {
		if entity, ok := v[module]; ok {
			return entity, ok
		}
	}
	return nil, false
}
