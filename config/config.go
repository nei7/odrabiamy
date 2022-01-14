package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Headers map[string]string `yaml:"headers"`
	Token   string            `yaml:"token"`
	Path    string            `yaml:"path"`
}

var Config config

func init() {
	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&Config); err != nil {
		log.Fatal(err)
	}

}
