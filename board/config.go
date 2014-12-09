package board

import (
	"ioutil"
	"encoding/json"
)

type ServerConfig struct {
	Redis struct {
		Address string `json:"address"`
		DB int `json:"db"`
		Auth string `json:"auth"`
	} `json:"redis"`,
	SecureSession struct {
		HashKey string `json:"hash_key"`,
		BlockKey string `json:"block_key"`
	} `json:"secure_session"`
}

func LoadConfig(path string) (*ServerConfig, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
	var config ServerConfig
    err = json.Unmarshal(data, &config)
    if err != nil {
		return nil, err
    }
    return &config, nil
}
