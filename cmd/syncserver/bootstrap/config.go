package bootstrap

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port     int    `yaml:"port"`
	FilePath string `yaml:"file-path"`
}

func LoadConfig(cfgPath string) (*ServerConfig, error) {
	cfgFile, err := os.Open(cfgPath)
	defer func() {
		if err = cfgFile.Close(); err != nil {
			return
		}
	}()
	if err != nil {
		return nil, err
	}
	cfg := &ServerConfig{}
	if err = yaml.NewDecoder(cfgFile).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
