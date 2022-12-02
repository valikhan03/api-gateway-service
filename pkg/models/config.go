package models

var ConfigGlobal Config

type Config struct {
	AuthService     AuthServiceConfigs     `yaml:"auth"`
	AuctionsService AuctionsServiceConfigs `yaml:"auctions"`
	CommandService  CommandServiceConfigs  `yaml:"command"`
}

type AuthServiceConfigs struct {
	URL string `yaml:"url"`
}

type AuctionsServiceConfigs struct {
}

type CommandServiceConfigs struct {
}
