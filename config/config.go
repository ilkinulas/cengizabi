package config

import "github.com/BurntSushi/toml"

type Config struct {
	BotApiToken string
	Podcast     Podcast
}

type Podcast struct {
	BaseUrl string
}

func Load(f string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(f, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
