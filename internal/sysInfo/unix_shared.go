//go:build unix

package sysinfo

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func (s *SystemInfo) GetHost() {
	host, err := os.Hostname()
	if err != nil {
		s.Host = "Unknown"
	}
	s.Host = host
}

func (s *SystemInfo) GetUser() {
	user, err := user.Current()
	if err != nil {
		s.User = "Unknown"
		return
	}
	s.User = user.Username
}

func (s *SystemInfo) GetTerm() {
	cmd := os.Getenv("TERM")
	if cmd == "" {
		s.Terminal = "Unknown"
		return
	}
	var terminal string
	switch {
	case strings.Contains(cmd, "tmux"):
		terminal = "tmux"
	case cmd == "xterm-256color":
		terminal = "xterm"
	case strings.Contains(cmd, "-"):
		var stringTerm []string
		stringTerm = strings.Split(cmd, "-")
		terminal = stringTerm[1]
	default:
		terminal = cmd
	}
	s.Terminal = terminal
}

func (s *SystemInfo) GetShell() {
	cmd := os.Getenv("SHELL")
	if cmd == "" {
		s.Shell = "Unknown"
		return
	}
	shellComp := strings.Split(cmd, "/")
	s.Shell = shellComp[len(shellComp)-1]
}

func (s *SystemInfo) GetDisk() {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		s.Disk = "Unknown"
		return
	}
	totalBytes := stat.Blocks * uint64(stat.Bsize)
	usedBytes := totalBytes - (stat.Bfree * uint64(stat.Bsize))
	const bToGb = 1073741824.0
	diskTotal := float64(totalBytes) / bToGb
	diskUsed := float64(usedBytes) / bToGb
	diskPorc := (diskUsed / diskTotal) * 100
	s.Disk = fmt.Sprintf("%.2f / %.2f (used: %.2f%%)", diskUsed, diskTotal, diskPorc)
}

func (s *SystemInfo) GetArch() {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		s.Arch = "Unknown"
		return
	}
	var arch []byte
	for _, b := range uname.Machine {
		if b == 0 {
			break
		}
		arch = append(arch, byte(b))
	}
	s.Arch = string(arch)
}
