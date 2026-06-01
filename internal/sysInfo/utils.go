package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ExecCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	out, err := cmd.Output()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}

func ReadFile(file string) string {
	doc, err := os.ReadFile(file)
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(doc))
}

func FormatUptime(uptimeSeconds string) string {
	seconds, err := strconv.Atoi(uptimeSeconds)
	if err != nil {
		return "Unknown"
	}
	d := seconds / 86400
	h := (seconds % 86400) / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60
	return fmt.Sprintf("%dd %dh %dm %ds", d, h, m, s)
}
