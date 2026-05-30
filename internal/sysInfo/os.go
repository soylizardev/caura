package sysinfo

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func (i *InfoOs) GetHost() {
	output, err := os.ReadFile("/proc/sys/kernel/hostname")
	if err != nil {
		i.Host = "Unknown"
		return
	}
	i.Host = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetOS() {
	output, err := os.ReadFile("/proc/sys/kernel/ostype")
	if err != nil {
		i.OS = "Unknown"
		return
	}
	i.OS = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetKernel() {
	output, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		i.Kernel = "Unknown"
		return
	}
	i.Kernel = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetUser() {
	output, err := user.Current()
	if err != nil {
		i.User = "Unknown"
		return
	}
	i.User = output.Username
}

func (i *InfoOs) GetTerm() {
	cmd := os.Getenv("TERM")
	if cmd == "" {
		i.Terminal = "Unknown"
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
	i.Terminal = terminal
}

func (i *InfoOs) GetShell() {
	cmd := os.Getenv("SHELL")
	if cmd == "" {
		i.Shell = "Unknown"
		return
	}
	shellComp := strings.Split(cmd, "/")
	i.Shell = shellComp[len(shellComp)-1]
}

func (i *InfoOs) GetTime() {
	read, err := os.ReadFile("/proc/uptime")

	if err != nil {
		i.Uptime = "Unknown"
		return
	}
	readStr := string(read)
	parts := strings.Split(readStr, " ")
	if len(parts) == 0 {
		i.Uptime = "Unknown"
		return
	}
	realsec := strings.Split(parts[0], ".")
	time := realsec[0]

	seconds, errParse := strconv.Atoi(time)
	if errParse != nil {
		i.Uptime = "Unknown"
		return
	}

	d := seconds / 86400

	h := (seconds % 86400) / 3600

	m := (seconds % 3600) / 60

	s := seconds % 60

	i.Uptime = fmt.Sprintf("%dd %dh %dm %ds", d, h, m, s)
}
