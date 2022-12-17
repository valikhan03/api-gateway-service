package models

var ConfigGlobal Config

type Config struct {
	AuthService     AuthServiceConfigs     `yaml:"auth_service"`
	AuctionsService AuctionsServiceConfigs `yaml:"auctions_service"`
	CommandService  CommandServiceConfigs  `yaml:"command_service"`
	SearchService   SearchServiceConfigs   `yaml:"search_service"`
}

type AuthServiceConfigs struct {
	URL string `yaml:"url"`
}

type AuctionsServiceConfigs struct {
	URL string `yaml:"url"`
}

type CommandServiceConfigs struct {
	URL string `yaml:"url"`
}

type SearchServiceConfigs struct {
	URL string `yaml:"url"`
}
