package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Conf struct {
	PkKey           string `yaml:"pk_key"`
	ContractAddress string `yaml:"contract_address"`
	ProviderUrl     string `yaml:"provider_url"`
}

func Parse() (Conf, error) {
	data, err := os.ReadFile("conf.yaml")
	if err != nil {
		return Conf{}, err
	}

	var config Conf

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Conf{}, err
	}
	return config, nil
}
