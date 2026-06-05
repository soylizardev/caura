package render

import "strings"

func defaultPadSep() int {
	return 2
}

func fieldWidth(fields []string) int {
	n := 0
	for _, f := range fields {
		if len(f) > n {
			n = len(f)
		}
	}
	return n
}

func sepWidth(sep string) int {
	return len(sep)
}

func padField(field string, padKey *int, maxFieldLen int) string {
	if padKey == nil || *padKey == 0 {
		return field
	}
	w := *padKey
	if w < maxFieldLen {
		w = maxFieldLen
	}
	return field + strings.Repeat(" ", w-len(field))
}

func gapSpaces(padSep *int, maxCombined, fieldLen, sepLen int) string {
	if padSep != nil && *padSep == 0 {
		return ""
	}
	gap := defaultPadSep()
	if padSep != nil {
		gap = *padSep
	}
	total := maxCombined + gap
	n := total - fieldLen - sepLen
	if n < 0 {
		n = 0
	}
	return strings.Repeat(" ", n)
}
