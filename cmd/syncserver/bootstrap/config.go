package bootstrap

import (
	"os"

	"github.com/kamijoucen/notesync/pkg/config"
	"gopkg.in/yaml.v3"
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
