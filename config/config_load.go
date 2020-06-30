package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// LoadConfig reads the config yaml file and returns the config
func LoadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open config file. Error: %v\n", err)
		return nil, err
	}
	defer f.Close()

	var c Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Printf("Failed to open config file. Error: %v\n", err)
		return nil, err
	}
	return &c, nil
}
