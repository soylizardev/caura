package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func DefaultConfig() *Config {
	return &Config{
		Header: Header{
			Enabled:   true,
			Separator: "@",
			SepColor:  "#6c7086",
			Fields:    []string{"User", "Host"},
		},
		Groups: []Group{
			{
				Title:      "   ------------------OS--------------------",
				TitleColor: "#cb6ff6",
				Separator:  ": ",
				SepColor:   "#6c7086",
				KeyColor:   "#b4befe",
				Fields:     []string{"OS", "Kernel", "Uptime", "Shell", "Terminal", "IP"},
			},
			{
				Title:      "   ------------------PC--------------------",
				TitleColor: "#cb6ff6",
				Separator:  ": ",
				SepColor:   "#6c7086",
				KeyColor:   "#b4befe",
				Fields:     []string{"PC", "CPU", "Arch", "Graphics", "Disk", "Memory", "Swap"},
			},
		},
		Footer: Footer{
			Enabled: true,
			Text:    "   ----------------------------------------",
			Color:   "#6c7086",
		},
	}
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			cfg := DefaultConfig()
			return cfg, writeDefault(path, cfg)
		}
		return nil, err
	}
	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func writeDefault(path string, cfg *Config) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
