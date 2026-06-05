package config

type Config struct {
	Header Header  `toml:"header"`
	Footer Footer  `toml:"footer"`
	Groups []Group `toml:"groups"`
}

type TextLine struct {
	Text  string `toml:"text"`
	Color string `toml:"color"`
}

type Header struct {
	Enabled    bool       `toml:"enabled"`
	Texts      []TextLine `toml:"texts"`
	Separator  string     `toml:"separator"`
	SepColor   string     `toml:"sep_color"`
	KeyColor   string     `toml:"key_color"`
	ValueColor string     `toml:"value_color"`
	PadKey     *int       `toml:"pad_key,omitempty"`
	PadSep     *int       `toml:"pad_sep,omitempty"`
	Fields     []string   `toml:"fields"`
}

type Footer struct {
	Enabled bool       `toml:"enabled"`
	Texts   []TextLine `toml:"texts"`
}

type Group struct {
	Title      string   `toml:"title"`
	TitleColor string   `toml:"title_color"`
	Separator  string   `toml:"separator"`
	SepColor   string   `toml:"sep_color"`
	KeyColor   string   `toml:"key_color"`
	ValueColor string   `toml:"value_color"`
	PadKey     *int     `toml:"pad_key,omitempty"`
	PadSep     *int     `toml:"pad_sep,omitempty"`
	Fields     []string `toml:"fields"`
}
