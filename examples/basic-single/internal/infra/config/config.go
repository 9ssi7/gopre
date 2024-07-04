package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	HttpPort    string `yaml:"http_port"`
	Database    struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SslMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	RunMigration bool `yaml:"run_migration"`
	RunSeed      bool `yaml:"run_seed"`
}

var configs *Config
var Path = "./config.yaml"

func ReadValue() *Config {
	if configs != nil {
		return configs
	}
	filename, _ := filepath.Abs(Path)
	cleanedDst := filepath.Clean(filename)
	yamlFile, _ := os.ReadFile(cleanedDst)
	err := yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		log.Fatal("error loading config.yaml ", err)
	}
	return configs
}
