package main

import "gopry/internal/sysInfo"

func main() {
	report := &sysinfo.SystemReport{}

	report.Render()
}
