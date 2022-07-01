package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Feature struct {
	Name    string `yaml:"name"`
	DSL     string `yaml:"dsl_location"`
	Version string `yaml:"version"`
}

type Config struct {
	Features []*Feature `yaml:"features"`
}

var (
	instance Config
)

func Load() {
	ymlFile := "./configs/features.yml"
	f, err := os.Open(ymlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&instance)
	if err != nil {
		panic(err)
	}
}

func GetCurrentConfig() *Config {
	return &instance
}
