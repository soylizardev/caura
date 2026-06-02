//go:build unix

package sysinfo

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"runtime"
	"strings"
	"syscall"
)

func (s *SystemInfo) GetHost() {
	host, err := os.Hostname()
	if err != nil {
		s.Host = "Unknown"
		return
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

func (s *SystemInfo) GetIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		s.IP = "Unknown"
		return
	}
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() || ipnet.IP.To4() == nil {
			continue
		}
		s.IP = ipnet.IP.String()
		return
	}
	s.IP = "Unknown"
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
	s.Arch = runtime.GOARCH
}
