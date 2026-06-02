package main

import "caura/internal/sysInfo"

func main() {
	report := &sysinfo.SystemInfo{}

	report.Render()
}
