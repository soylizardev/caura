package main

import "github.com/soylizardev/caura/internal/sysInfo"

func main() {
	report := &sysinfo.SystemInfo{}

	report.Render()
}
