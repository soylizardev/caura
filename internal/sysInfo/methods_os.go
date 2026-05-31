package sysinfo

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func (i *InfoOs) GetHost(file string) {
	output, err := os.ReadFile(file)
	if err != nil {
		i.Host = "Unknown"
		return
	}
	i.Host = strings.TrimSpace(string(output))
}

func (i *InfoOs) GetOS(file1, file2 string) {
	content, errDis := os.ReadFile(file1)
	var distro string
	if errDis != nil {
		distro = "Unknown"
	}

	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			div := strings.Split(line, "=")
			if len(div) > 1 {
				distro = strings.Trim(strings.TrimSpace(div[1]), "\"")
				break
			}
		}
	}

	output, err := os.ReadFile(file2)
	if err != nil {
		i.OS = "Unknown"
		return
	}
	i.OS = fmt.Sprintf("%s (%s)", distro, strings.TrimSpace(string(output)))
}

func (i *InfoOs) GetKernel(file string) {
	output, err := os.ReadFile(file)
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

func (i *InfoOs) GetTerm(envVar string) {
	cmd := os.Getenv(envVar)
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

func (i *InfoOs) GetShell(envVar string) {
	cmd := os.Getenv(envVar)
	if cmd == "" {
		i.Shell = "Unknown"
		return
	}
	shellComp := strings.Split(cmd, "/")
	i.Shell = shellComp[len(shellComp)-1]
}

func (i *InfoOs) GetUptime(file string) {
	read, err := os.ReadFile(file)

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
