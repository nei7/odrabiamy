package config

import (
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type config struct {
	Headers map[string]string `yaml:"headers"`
	Token   string            `yaml:"token"`
}

var Config config

func init() {
	f, err := os.Open("./config.yaml")
	if err != nil {
		zap.S().Fatal(err)
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&Config); err != nil {
		zap.S().Fatal(err)
	}

}
