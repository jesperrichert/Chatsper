package file

import (
	"zip.jespersen.chatsper/shared/database"
)

type Config struct {
	TwitchPlatform    TwitchPlatform `yaml:"TwitchPlatform"`
	Database          Database       `yaml:"Database"`
	Version           string         `yaml:"CONFIG_VERSION"`
	DisableAutoUpdate bool           `yaml:"DisableAutoUpdate"`
}

type TwitchPlatform struct {
	BotAccountUserId  uint   `yaml:"BotAccountUserId"`
	ApplicationId     string `yaml:"ApplicationId"`
	ApplicationSecret string `yaml:"ApplicationSecret"`
	RedirectUrl       string `yaml:"RedirectUrl"`
}

type Database struct {
	Type          database.DatabaseType `yaml:"Type"`
	ConnectionUrl string                `yaml:"ConnectionUrl"`
}
