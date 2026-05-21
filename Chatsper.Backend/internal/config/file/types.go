package file

import (
	"zip.jespersen.chatsper/Chatsper.Backend/shared/database"
)

type Config struct {
	Bot      Bot      `yaml:"Bot"`
	Database Database `yaml:"Database"`
	Version  string   `yaml:"CONFIG_VERSION"`
}

type Bot struct {
	UserId string `yaml:"UserId"`
	ID     string `yaml:"ID"`
	Secret string `yaml:"Secret"`
}

type Database struct {
	Type          database.DatabaseType `yaml:"Type"`
	ConnectionUrl string                `yaml:"ConnectionUrl"`
}
