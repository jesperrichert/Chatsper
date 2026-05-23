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
	ApplicationId     string `yaml:"ApplicationId"`
	ApplicationSecret string `yaml:"ApplicationSecret"`
}

type Database struct {
	Type          database.DatabaseType `yaml:"Type"`
	ConnectionUrl string                `yaml:"ConnectionUrl"`
}
