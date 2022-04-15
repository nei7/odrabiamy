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
	S3      struct {
		Bucket string `yaml:"bucket"`
		Region string `yaml:"region"`
	} `yaml:"s3"`
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
