package main

import "gopry/internal/sysInfo"

func main() {
	report := &sysinfo.SystemInfo{}

	report.Render()
}
