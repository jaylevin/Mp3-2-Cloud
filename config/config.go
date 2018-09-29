package config

import (
	"os"
	"encoding/json"
)

type Config struct {
	AccessKey  string  // from digital oceans control panel
	SecretKey  string  // from digital oceans control panel
	Endpoint   string  // nyc3.digitaloceanspaces.com
	BucketName string  // the name of the digital ocean's space (on digital oceans control panel)
}

func DefaultConfig() *Config {
	return &Config {
		AccessKey: "your_access_key",
		SecretKey: "your_secret_key",
		Endpoint:  "nyc3.digitaloceanspaces.com",
		BucketName: "dev",
	}
}

func Load() (*Config, error) {
	f, err := os.Open("config.json")
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer f.Close()

	conf := new(Config)
	err = json.NewDecoder(f).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
