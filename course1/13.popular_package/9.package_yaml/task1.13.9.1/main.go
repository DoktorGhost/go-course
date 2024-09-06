package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func getYAML(configs []Config) (string, error) {

	encoder, err := yaml.Marshal(configs)
	if err != nil {
		return "", err
	}

	return string(encoder), nil
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
	configs := []Config{
		{
			Server: Server{
				Port: "8080",
			},
			Db: Db{
				Host:     "localhost",
				Port:     "5432",
				User:     "postgres",
				Password: "postgresPass",
			},
		},
		{
			Server: Server{
				Port: "8081",
			},
			Db: Db{
				Host:     "localhost",
				Port:     "5433",
				User:     "mongo",
				Password: "mongoPass",
			},
		},
	}

	str, err := getYAML(configs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(str)
}
