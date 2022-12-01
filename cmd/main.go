package main

import (
	"fmt"
	"github.com/Jeff-Moorhead-PLine/oauth-playground/internal/config"
	"log"
	"os"
)

func main() {

	configPath := os.Getenv("OAUTH_CONFIG_PATH")
	if configPath == "" {
		log.Fatal("no config path set. Please set OAUTH_CONFIG_PATH.")
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
