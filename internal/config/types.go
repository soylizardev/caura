package config

type Config struct {
	Header Header  `toml:"header"`
	Footer Footer  `toml:"footer"`
	Groups []Group `toml:"groups"`
}

type Header struct {
	Enabled   bool     `toml:"enabled"`
	Separator string   `toml:"separator"`
	Fields    []string `toml:"fields"`
}

type Footer struct {
	Enabled bool   `toml:"enabled"`
	Text    string `toml:"text"`
}

type Group struct {
	Title     string   `toml:"title"`
	Separator string   `toml:"separator"`
	Fields    []string `toml:"fields"`
}
