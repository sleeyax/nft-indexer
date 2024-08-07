package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	OpenSea   OpenSeaConfig
	Gcloud    GcloudConfig
	Alchemy   AlchemyConfig
	QuickNode QuickNodeConfig
	Moralis   ApiKeyConfig
	Mnemonic  ApiKeyConfig
	Zora      ApiKeyConfig
	Reservoir ApiKeyConfig
}

type OpenSeaConfig struct {
	ApiKeys []string
}

type GcloudConfig struct {
	Queue     ApiKeyConfig
	Firestore FirestoreConfig
}

type AlchemyConfig struct {
	JsonRpc struct {
		MainNet []string
		Goerli  []string
	}
}

type QuickNodeConfig struct {
	Wss struct {
		MainNet []string
	}
}

type FirestoreConfig struct {
	// Service account JSON string or path to creds.json
	ServiceAccount string
}

type InfuraConfig struct {
	Ipfs []InfuraIPFSConfig
}

type InfuraIPFSConfig struct {
	ProjectId string
	Secret    string
}

type ApiKeyConfig struct {
	ApiKey string
}

// ParseConfig loads and parses the specified configuration file.
func ParseConfig(fileName string) (*Configuration, error) {
	if split := strings.Split(fileName, "."); len(split) > 0 {
		fileName = split[0]
	}

	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
