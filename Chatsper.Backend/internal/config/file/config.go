package file

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/pterm/pterm"
	Log "zip.jespersen.chatsper/internal/utils/log"
	"zip.jespersen.chatsper/shared/database"
	"zip.jespersen.chatsper/shared/static"
)

var ConfigTemplate = &Config{
	TwitchPlatform: TwitchPlatform{
		ApplicationId:     "",
		ApplicationSecret: "",
		RedirectUrl:       "http://localhost:3000/platform/twitch/auth/callback",
	},
	Database: Database{
		Type:          database.SqliteDataBase,
		ConnectionUrl: "postgres://postgres:postgres@localhost:5432/postgres",
	},
	Version: static.ConfigVersion,
}

func LoadFromFile(file string) (*Config, error) {
	Log.Info("Loading config from file: " + file)
	data, err := os.ReadFile(file)
	if err != nil {
		Log.Error("No Config found for Chatsper... Please open config.yml and configure your Chatsper.", false)
		Log.Debug("Then Restart the Bot...")

		template := ConfigTemplate
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
		Log.Error("", true)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)

	if static.ConfigVersion != config.Version {
		pterm.DefaultBasicText.Println(pterm.Red("Config version mismatch. ") + "Please" + pterm.LightBlue(" backup ") + "your config and lets" + pterm.LightMagenta(" Chatsper ") + "regenerate the config!")
		os.Exit(0)
	}

	if err != nil {
		return nil, err
	}
	return &config, nil
}

func SaveToFile(file string, config *Config) error {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(file)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
