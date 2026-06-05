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
			Fields:    []string{"User", "Hostname"},
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
	data := `# =============================================================================
# caura v0.2.0 — Configuration file
# =============================================================================
# This file is auto-generated on first run.
# Edit it to customize the output of caura.
#
# Colors: #RRGGBB format (e.g., #ff0000 for red).
#         Leave empty for terminal default color.
# Fields: see each section for available field names.
# =============================================================================

# --- Header ----------------------------------------------------------------
# The first line of output (e.g., user@hostname).
[header]
enabled = true

# Optional text line printed before the fields.
# text = ""
# color = ""

# Separator between header fields.
separator = "@"
sep_color = "#6c7086"

# Color for field values in the header.
# value_color = ""

# Available fields: User, Hostname
# User     -> current username
# Hostname -> system hostname
fields = ["User", "Hostname"]

# --- Groups ----------------------------------------------------------------
# Each group is a section with a title and fields.
# You can add, remove, or reorder groups freely.
# Available fields for any group:
#   OS, Kernel, Uptime, Shell, Terminal, IP,
#   PC, CPU, Arch, Graphics, Disk, Memory, Swap,
#   Hostname, Host
#   Host     -> full PC description (product name/model)
#   Hostname -> short system hostname
#   (yes, Host != Hostname; Host is the product, Hostname is the network name)

[[groups]]
title = "   ------------------OS--------------------"
title_color = "#cb6ff6"
separator = ": "
sep_color = "#6c7086"
key_color = "#b4befe"
# value_color = ""
fields = ["OS", "Kernel", "Uptime", "Shell", "Terminal", "IP"]

[[groups]]
title = "   ------------------PC--------------------"
title_color = "#cb6ff6"
separator = ": "
sep_color = "#6c7086"
key_color = "#b4befe"
# value_color = ""
fields = ["PC", "CPU", "Arch", "Graphics", "Disk", "Memory", "Swap"]

# --- Footer ----------------------------------------------------------------
# A closing line at the end of the output.
[footer]
enabled = true
text = "   ----------------------------------------"
color = "#6c7086"
`
	return os.WriteFile(path, []byte(data), 0644)
}
