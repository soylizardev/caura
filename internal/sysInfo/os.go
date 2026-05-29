package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func (i *InfoOs) GetHost() {
	cmd := exec.Command("uname", "-n")
	output, err := cmd.Output()
	if err != nil {
		i.Host = fmt.Sprintln("Error:", err)
		return
	}
	i.Host = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetOS() {
	cmd := exec.Command("uname", "-o")
	output, err := cmd.Output()
	if err != nil {
		i.OS = fmt.Sprintln("Error:", err)
		return
	}
	i.OS = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetKernel() {
	cmd := exec.Command("uname", "-rs")
	output, err := cmd.Output()
	if err != nil {
		i.Kernel = fmt.Sprintln("Error:", err)
		return
	}
	i.Kernel = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetUser() {
	cmd := exec.Command("whoami")
	output, err := cmd.Output()
	if err != nil {
		i.User = fmt.Sprintln("Error", err)
		return
	}
	i.User = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetTerm() {
	cmd := os.Getenv("TERM")
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
	com := exec.Command(terminal, "--version")
	termVer, _ := com.Output()
	i.Terminal = fmt.Sprintf("%s - %s", terminal, termVer)
}

func (i *InfoOs) GetShell() {
	cmd := os.Getenv("SHELL")
	shellComp := strings.Split(cmd, "/")
	i.Shell = shellComp[len(shellComp)-1]
}

func (i *InfoOs) GetTime() {
	read, err := os.ReadFile("/proc/uptime")

	readStr := string(read)
	parts := strings.Split(readStr, " ")
	realsec := strings.Split(parts[0], ".")
	time := realsec[0]

	if err != nil {
		i.Uptime = err.Error()
	}

	seconds, _ := strconv.Atoi(time)

	d := seconds / 86400

	h := (seconds % 86400) / 3600

	m := (seconds % 3600) / 60

	s := seconds % 60

	i.Uptime = fmt.Sprintf("%dd %dh %dm %ds", d, h, m, s)
}
