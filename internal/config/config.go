package config

import (
	"os"
	"path/filepath"
	"strings"

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
				Title:      "------------------OS--------------------",
				TitleColor: "#cb6ff6",
				Separator:  ": ",
				SepColor:   "#6c7086",
				KeyColor:   "#b4befe",
				Fields:     []string{"OS", "Kernel", "Uptime", "Shell", "Terminal", "IP"},
			},
			{
				Title:      "------------------PC--------------------",
				TitleColor: "#cb6ff6",
				Separator:  ": ",
				SepColor:   "#6c7086",
				KeyColor:   "#b4befe",
				Fields:     []string{"PC", "CPU", "Arch", "Graphics", "Disk", "Memory", "Swap"},
			},
		},
		Footer: Footer{
			Enabled: true,
			Texts: []TextLine{
				{Text: "----------------------------------------", Color: "#6c7086"},
			},
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

	if cfg.Theme.Enabled && cfg.Theme.Path != "" {
		themePath := expandPath(cfg.Theme.Path)
		themeData, err := os.ReadFile(themePath)
		if err == nil {
			var themeCfg Config
			if err := toml.Unmarshal(themeData, &themeCfg); err == nil {
				return &themeCfg, nil
			}
		}
	}

	return &cfg, nil
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(home, path[2:])
		}
	}
	return path
}

func writeDefault(path string, cfg *Config) error {
	baseDir := filepath.Dir(path)
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return err
	}
	os.MkdirAll(filepath.Join(baseDir, "theme"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "logo", "text"), 0755)
	data := `# =============================================================================
# caura v0.2.2 — Configuration file
# =============================================================================
# This file is auto-generated on first run.
# Edit it to customize the output of caura.
#
# Colors: #RRGGBB format (e.g., #ff0000 for red).
#         Leave empty for terminal default color.
# Fields: see each section for available field names.
# =============================================================================

# --- Theme ------------------------------------------------------------------
# If enabled and the path points to an existing TOML file, the entire
# configuration is loaded from that file (the rest of this file is ignored).
# If disabled, or the path is invalid, this file is used as-is.
# [theme]
# enabled = true
# path = "~/.config/caura/theme/mi_tema.toml"

# --- Header ----------------------------------------------------------------
# The first line of output (e.g., user@hostname).
[header]
enabled = true

# Optional text lines printed before the fields.
# Add as many as you want:
# [[header.texts]]
# text = "My ASCII logo"
# color = "#ffffff"

# Separator between header fields.
separator = "@"
sep_color = "#6c7086"

# Color for field values in the header.
# value_color = ""

# Padding: pad_sep controls spaces between separator and next value.
#   Not set / 0 → no extra gap.
#   > 0         → N spaces after separator.
# pad_sep = 2

# Available fields: User, Hostname
# User     -> current username
# Hostname -> system hostname
fields = ["User", "Hostname"]

# --- Groups ----------------------------------------------------------------
# Each group is a section with a title and fields.
# You can add, remove, or reorder groups freely.
#
# Available fields for any group:
#   OS, Kernel, Uptime, Shell, Terminal, IP,
#   PC, CPU, Arch, Graphics, Disk, Memory, Swap,
#   Hostname, Host
#   Host     -> full PC description (product name/model)
#   Hostname -> short system hostname
#   (yes, Host != Hostname; Host is the product, Hostname is the network name)
#
# Padding controls alignment per group:
#   pad_key   Not set / 0 → key natural width.
#             > 0         → key padded to exactly N chars.
#   pad_sep   Not set     → all values aligned (keys padded to longest in group).
#             0           → value glued to separator.
#             > 0         → N fixed spaces between separator and value.

[[groups]]
title = "------------------OS--------------------"
title_color = "#cb6ff6"
separator = ": "
sep_color = "#6c7086"
key_color = "#b4befe"
# value_color = ""
# pad_key = 0
# pad_sep = 2
fields = ["OS", "Kernel", "Uptime", "Shell", "Terminal", "IP"]

[[groups]]
title = "------------------PC--------------------"
title_color = "#cb6ff6"
separator = ": "
sep_color = "#6c7086"
key_color = "#b4befe"
# value_color = ""
# pad_key = 0
# pad_sep = 2
fields = ["PC", "CPU", "Arch", "Graphics", "Disk", "Memory", "Swap"]

# --- Footer ----------------------------------------------------------------
# A closing line at the end of the output.
# Add as many text lines as you want:
# [[footer.texts]]
# text = "My footer text"
# color = "#6c7086"
[footer]
enabled = true

[[footer.texts]]
text = "----------------------------------------"
color = "#6c7086"
`
	return os.WriteFile(path, []byte(data), 0644)
}
