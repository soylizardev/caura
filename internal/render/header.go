package render

import (
	"fmt"
	"strings"

	"github.com/soylizardev/caura/internal/config"
	"github.com/soylizardev/caura/internal/sysInfo"
)

func renderHeader(s *sysinfo.SystemInfo, h config.Header) {
	for _, t := range h.Texts {
		for _, line := range strings.Split(t.Text, "\n") {
			fmt.Println(colorize(t.Color, "   "+line))
		}
	}
	if len(h.Fields) == 0 {
		return
	}

	fmt.Print("   ")
	for i, field := range h.Fields {
		if i > 0 {
			gap := ""
			if h.PadSep != nil && *h.PadSep > 0 {
				gap = strings.Repeat(" ", *h.PadSep)
			}
			fmt.Printf("%s%s",
				colorize(h.SepColor, h.Separator),
				gap,
			)
			fmt.Print(colorize(h.ValueColor, getFieldValue(s, field)))
		} else {
			fmt.Print(colorize(h.KeyColor, getFieldValue(s, field)))
		}
	}
	fmt.Println()
}
