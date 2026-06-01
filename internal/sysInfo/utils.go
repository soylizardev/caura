package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
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

func UserCurrent() string {
	output, err := user.Current()
	if err != nil {
		return "Unknown"
	}
	return output.Username
}

func EnvTerm(envVar string) string {
	cmd := os.Getenv(envVar)
	if cmd == "" {
		return "Unknown"
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
	return terminal
}

func EnvShell(envVar string) string {
	cmd := os.Getenv(envVar)
	if cmd == "" {
		return "Unknown"
	}
	shellComp := strings.Split(cmd, "/")
	return shellComp[len(shellComp)-1]
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
