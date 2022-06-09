package nft_indexer

import "github.com/spf13/viper"

type Configuration struct {
	OpenSea   OpenSeaConfig
	Gcloud    GcloudConfig
	Alchemy   AlchemyConfig
	Moralis   MoralisConfig
	Mnemonic  MnemonicConfig
	QuickNode QuickNodeConfig
}

type OpenSeaConfig struct {
	ApiKeys []string
}

type GcloudConfig struct {
	Queue struct {
		ApiKey string
	}
}

type AlchemyConfig struct {
	JsonRpc struct {
		MainNet []string
		Goerli  []string
	}
}

type MoralisConfig struct {
	ApiKey string
}

type MnemonicConfig struct {
	ApiKey string
}

type QuickNodeConfig struct {
	Wss struct {
		MainNet []string
	}
}

type InfuraConfig struct {
	Ipfs []InfuraIPFSConfig
}

type InfuraIPFSConfig struct {
	ProjectId string
	Secret    string
}

// Configure loads and parses the configuration file.
// It should be stored relative to the working directory.
func Configure() (*Configuration, error) {
	viper.SetConfigName("config") // name of config file (without extension)
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
