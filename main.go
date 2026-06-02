package main

import (
	"flag"
	"fmt"
	"github.com/soylizardev/caura/internal/sysInfo"
)

const version = "v0.1.3"

func main() {
	showVersion := flag.Bool("version", false, "Show version")
	flag.BoolVar(showVersion, "v", false, "Show version (short)")
	flag.Parse()

	if *showVersion {
		fmt.Println("caura", version)
		return
	}

	report := &sysinfo.SystemInfo{}
	report.Render()
}
