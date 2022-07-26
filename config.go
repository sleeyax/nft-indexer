package nft_indexer

import (
	"github.com/spf13/viper"
	"strings"
)

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
	Firestore FirestoreConfig
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

type FirestoreConfig struct {
	ServiceAccount FirestoreServiceAccount
}

type FirestoreServiceAccount struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

type InfuraConfig struct {
	Ipfs []InfuraIPFSConfig
}

type InfuraIPFSConfig struct {
	ProjectId string
	Secret    string
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
