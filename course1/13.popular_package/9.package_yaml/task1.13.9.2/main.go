package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func getConfigFromYAML(data []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func main() {
	data := []byte(`
server:
  port: "8080"
db:
  host: localhost
  port: "5432"
  user: postgres
  password: postgresPass
`)

	config, err := getConfigFromYAML(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
