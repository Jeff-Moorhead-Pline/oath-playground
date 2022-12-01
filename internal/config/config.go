package config

import (
	"encoding/json"
	"os"
)

type config struct {
	AuthorizationUri string `json:"authorization_uri"`
	UserInfoUri      string `json:"user_info_uri"`
	TokenUri         string `json:"token_uri"`
	JwksUri          string `json:"jwks_uri"`
	ClientId         string `json:"client_id"`
	ClientSecret     string `json:"client_secret"`
}

func Load(path string) (*config, error) {

	cnt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg config
	if err := json.Unmarshal(cnt, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
