package config

type ServerConfig struct {
	Port     int    `yaml:"port"`
	FilePath string `yaml:"file-path"`
}
