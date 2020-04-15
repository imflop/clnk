package clnkserver

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config clnkserver ...
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
	OriginURL struct {
		Frontend string `yaml:"frontend"`
	} `yaml:"origin_url"`
}

// NewConfig ...
func NewConfig(configPath string) (*Config, error) {
	// create config struct
	config := &Config{}

	// validate path
	if err := validateConfigPath(configPath); err != nil {
		return nil, err
	}

	// open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// init yaml decoder
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
