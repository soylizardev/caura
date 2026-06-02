//go:build freebsd

package sysinfo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (s *SystemInfo) GetOS() {
	nameOs := ExecCommand("sysctl", "-n", "kern.ostype")
	versOs := ExecCommand("sysctl", "-n", "kern.osrelease")
	s.OS = fmt.Sprintf("%s (%s)", nameOs, versOs)
}

func (s *SystemInfo) GetKernel() {
	cmd := ExecCommand("sysctl", "-n", "kern.version")
	fields := strings.Fields(cmd)
	if len(fields) > 1 {
		s.Kernel = fields[1]
	} else {
		s.Kernel = "Unknown"
	}
}

func (s *SystemInfo) GetUptime() {
	cmd := ExecCommand("sysctl", "-n", "kern.boottime")

	parts := strings.Fields(cmd)
	var sec string
	if len(parts) > 1 {
		sec = strings.TrimRight(parts[3], ",")
		secnum, _ := strconv.Atoi(sec)
		num := time.Now().Unix() - int64(secnum)
		s.Uptime = FormatUptime(strconv.FormatInt(num, 10))
	} else {
		s.Uptime = "Unknown"
	}
}
