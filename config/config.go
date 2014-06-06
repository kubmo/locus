package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiKey string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

func FromFile(path string) (c Config, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &c)
	return
}
