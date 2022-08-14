package models


var ConfigGlobal Config

type Config struct{
	AuthService AuthServiceConfigs `yaml:"auth"`
}

type AuthServiceConfigs struct{
	URL string `yaml:"url"`
}