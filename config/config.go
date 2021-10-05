package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ServerConfig struct {
	Api struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		ApiKey string `yaml:"api_key"`
	}
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBname   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var Server ServerConfig

func init() {
	LoadYamlConfig(&Server)
}

func LoadYamlConfig(config *ServerConfig) {
	filePath := "./Config/server.yaml"
	yamlFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Printf("open file %v error, message: %v", filePath, err)
	}

	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		log.Fatalf("parse yaml file fail, message: %v", err)
	}
}
