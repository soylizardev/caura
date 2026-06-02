//go:build linux

package sysinfo

import (
	"fmt"
	"strings"
)

func (s *SystemInfo) GetOS() {
	content := ReadFile("/etc/os-release")
	var distro string
	lines := strings.SplitSeq(content, "\n")
	for line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			div := strings.Split(line, "=")
			if len(div) > 1 {
				distro = strings.Trim(strings.TrimSpace(div[1]), "\"")
				break
			}
		}
	}

	output := ReadFile("/proc/sys/kernel/ostype")
	s.OS = fmt.Sprintf("%s (%s)", distro, strings.TrimSpace(string(output)))
}

func (s *SystemInfo) GetKernel() {
	output := ReadFile("/proc/sys/kernel/osrelease")
	s.Kernel = output
}

func (s *SystemInfo) GetUptime() {
	read := ReadFile("/proc/uptime")
	parts := strings.Fields(read)
	if len(parts) == 0 {
		s.Uptime = "Unknown"
		return
	}
	secs, _, _ := strings.Cut(parts[0], ".")

	s.Uptime = FormatUptime(secs)
}
