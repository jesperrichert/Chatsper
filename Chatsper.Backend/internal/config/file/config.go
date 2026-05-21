package file

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/pterm/pterm"
	"zip.jespersen.chatsper/Chatsper.Backend/shared/database"
)

var Version = "1.0.0"

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
			Database: Database{
				Type:          database.DatabaseType(0),
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
