package file

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/pterm/pterm"
	Log "zip.jespersen.chatsper/internal/utils"
	"zip.jespersen.chatsper/shared/database"
)

var Version = "1.0.2"

func LoadFromFile(file string) (*Config, error) {
	Log.Info("Loading config from file: " + file)
	data, err := os.ReadFile(file)
	if err != nil {
		Log.Error("No Config found for Chatsper... Please open config.yml and configure your Chatsper.", false)
		Log.Debug("Then Restart the Bot...")

		template := &Config{
			TwitchPlatform: TwitchPlatform{
				ApplicationId:     "",
				ApplicationSecret: "",
				RedirectUrl:       "http://localhost:3000/platform/twitch/auth/callback",
			},
			Database: Database{
				Type:          database.SqliteDataBase,
				ConnectionUrl: "postgres://postgres:postgres@localhost:5432/postgres",
			},
			Version: Version,
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
		Log.Error("", true)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)

	if Version != config.Version {
		pterm.DefaultBasicText.Println(pterm.Red("Config version mismatch. ") + "Please" + pterm.LightBlue(" backup ") + "your config and lets" + pterm.LightMagenta(" Chatsper ") + "regenerate the config!")
		os.Exit(0)
	}

	if err != nil {
		return nil, err
	}
	return &config, nil
}
