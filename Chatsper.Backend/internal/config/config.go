package config

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

var CONFIG_VERSION = "1.0.0"

func LoadFromFile(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("No config file found")

		template := &Config{
			Bot: Bot{
				ID:     "",
				Secret: "",
				UserId: "",
			},
			Version: CONFIG_VERSION,
		}
		bytes, err := yaml.Marshal(template)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Create(file)
		if err != nil {
			return nil, err
		}
		err = os.WriteFile(file, bytes, 0644)
		if err != nil {
			return nil, err
		}
	}
	var config Config
	err = yaml.Unmarshal(data, &config)

	if CONFIG_VERSION != config.Version {
		log.Fatal("Config version mismatch. Please backup your config and lets Chatsper regenerate the config!")
	}

	if err != nil {
		return nil, err
	}
	return &config, nil
}
