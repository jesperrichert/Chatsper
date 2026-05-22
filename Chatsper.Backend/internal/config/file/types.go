package file

import (
	"zip.jespersen.chatsper/shared/database"
)

type Config struct {
	TwitchPlatform TwitchPlatform `yaml:"TwitchPlatform"`
	Database       Database       `yaml:"Database"`
	Version        string         `yaml:"CONFIG_VERSION"`
}

type TwitchPlatform struct {
	ApplicationUserId string `yaml:"ApplicationUserId"`
	ApplicationID     string `yaml:"ApplicationID"`
	ApplicationSecret string `yaml:"ApplicationSecret"`
}

type Database struct {
	Type          database.DatabaseType `yaml:"Type"`
	ConnectionUrl string                `yaml:"ConnectionUrl"`
}
