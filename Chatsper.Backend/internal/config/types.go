package config

type Config struct {
	Bot     Bot    `yaml:"Bot"`
	Version string `yaml:"CONFIG_VERSION"`
}

type Bot struct {
	UserId string `yaml:"UserId"`
	ID     string `yaml:"ID"`
	Secret string `yaml:"Secret"`
}
