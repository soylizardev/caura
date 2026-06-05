package render

import (
	"fmt"
	"strconv"

	"github.com/soylizardev/caura/internal/config"
	"github.com/soylizardev/caura/internal/sysInfo"
)

func Render(s *sysinfo.SystemInfo, cfg *config.Config) {
	if cfg.Header.Enabled {
		if cfg.Header.Text != "" {
			fmt.Println(colorize(cfg.Header.Color, cfg.Header.Text))
		}
		if len(cfg.Header.Fields) > 0 {
			fmt.Print("   ")
			for i, field := range cfg.Header.Fields {
				if i > 0 {
					fmt.Print(colorize(cfg.Header.SepColor, cfg.Header.Separator))
				}
				fmt.Print(colorize(cfg.Header.ValueColor, getFieldValue(s, field)))
			}
			fmt.Println()
		}
	}
	for _, group := range cfg.Groups {
		if group.Title != "" {
			fmt.Println(colorize(group.TitleColor, group.Title))
		}
		for _, field := range group.Fields {
			value := getFieldValue(s, field)
			fmt.Printf("   %s%s%s\n",
				colorize(group.KeyColor, field),
				colorize(group.SepColor, group.Separator),
				colorize(group.ValueColor, value),
			)
		}
	}
	if cfg.Footer.Enabled {
		fmt.Println(colorize(cfg.Footer.Color, cfg.Footer.Text))
	}
}

func getFieldValue(s *sysinfo.SystemInfo, field string) string {
	switch field {
	case "User":
		return s.User
	case "Host":
		return s.Pc
	case "Hostname":
		return s.Host
	case "Terminal":
		return s.Terminal
	case "Shell":
		return s.Shell
	case "IP":
		return s.IP
	case "Arch":
		return s.Arch
	case "Disk":
		return s.Disk
	case "OS":
		return s.OS
	case "Kernel":
		return s.Kernel
	case "Uptime":
		return s.Uptime
	case "PC":
		return s.Pc
	case "CPU":
		return s.CPU
	case "Graphics":
		return s.Graphic
	case "Memory":
		return s.Ram
	case "Swap":
		return s.Swap
	default:
		return "Unknown"
	}
}

func colorize(hex, text string) string {
	if hex == "" {
		return text
	}
	start := 0
	if hex[0] == '#' {
		start = 1
	}

	s := hex[start:]
	if len(s) != 6 {
		return text
	}

	r, err := strconv.ParseUint(s[0:2], 16, 8)
	if err != nil {
		return text
	}
	g, err := strconv.ParseUint(s[2:4], 16, 8)
	if err != nil {
		return text
	}
	b, err := strconv.ParseUint(s[4:6], 16, 8)
	if err != nil {
		return text
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, text)
}
