package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/kamijoucen/notesync/pkg/config"
)

func LoadConfig(cfgPath string) (*config.ServerConfig, error) {

	cfgFile, err := os.Open(cfgPath)
	defer func() {
		if err = cfgFile.Close(); err != nil {
			return
		}
	}()
	if err != nil {
		return nil, err
	}
	cfg := &config.ServerConfig{}
	if err = yaml.NewDecoder(cfgFile).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
