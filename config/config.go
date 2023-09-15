package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Conf struct {
	PkKey           string `yaml:"pk_key"`
	ContractAddress string `yaml:"contract_address"`
	ProviderHost    string `yaml:"provider_host"`
	ProviderPort    int    `yaml:"provider_port"`
}

func (c Conf) ProviderURL() string {
	return fmt.Sprintf("%s:%d", c.ProviderHost, c.ProviderPort)
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
