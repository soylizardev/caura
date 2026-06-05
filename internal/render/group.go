package render

import (
	"fmt"

	"github.com/soylizardev/caura/internal/config"
	"github.com/soylizardev/caura/internal/sysInfo"
)

func renderGroup(s *sysinfo.SystemInfo, g config.Group) {
	if g.Title != "" {
		fmt.Println(colorize(g.TitleColor, "   "+g.Title))
	}

	maxFieldLen := fieldWidth(g.Fields)
	maxCombined := maxFieldLen + sepWidth(g.Separator)

	for _, field := range g.Fields {
		value := getFieldValue(s, field)
		key := padField(field, g.PadKey, maxFieldLen)
		gap := gapSpaces(g.PadSep, maxCombined, len(field), sepWidth(g.Separator))
		fmt.Printf("   %s%s%s%s\n",
			colorize(g.KeyColor, key),
			colorize(g.SepColor, g.Separator),
			gap,
			colorize(g.ValueColor, value),
		)
	}
}
