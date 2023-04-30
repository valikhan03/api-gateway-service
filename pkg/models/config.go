package models

import(
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var ConfigGlobal Config

type Config struct {
	AuthService     AuthServiceConfigs     `yaml:"auth_service"`
	AuctionsService AuctionsServiceConfigs `yaml:"auctions_service"`
	CommandService  CommandServiceConfigs  `yaml:"command_service"`
	SearchService   SearchServiceConfigs   `yaml:"search_service"`
	QueryService    QueryService           `yaml:"query_service"`
	PaymentService  PaymentService         `yaml:"payment_service"`
}

func InitConfigs() {
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil{
		log.Fatalf("Unable to oper config-file: %v", err)
	}

	err = yaml.Unmarshal(file, &ConfigGlobal)
	if err != nil {
		log.Fatalf("Config file unmarhsal error: %v", err)
	}
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

type QueryService struct {
	URL string `yaml:"url"`
}

type PaymentService struct {
	URL string `yaml:"url"`
}
