package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/soylizardev/caura/internal/config"
	"github.com/soylizardev/caura/internal/render"
	"github.com/soylizardev/caura/internal/sysInfo"
)

const version = "v0.2.1"

func main() {
	showVersion := flag.Bool("version", false, "Show version")
	flag.BoolVar(showVersion, "v", false, "Show version (short)")
	flag.Parse()

	if *showVersion {
		fmt.Println("caura", version)
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error obteniendo home:", err)
		os.Exit(1)
	}
	cfg, err := config.Load(filepath.Join(home, ".config", "caura", "config.toml"))
	if err != nil {
		fmt.Println("Error cargando configuración:", err)
		os.Exit(1)
	}

	s := &sysinfo.SystemInfo{}
	s.GetUser()
	s.GetHost()
	s.GetOS()
	s.GetKernel()
	s.GetUptime()
	s.GetShell()
	s.GetTerm()
	s.GetIP()
	s.GetPc()
	s.GetCPU()
	s.GetArch()
	s.GetGraphic()
	s.GetDisk()
	s.GetRam()
	s.GetSwap()

	render.Render(s, cfg)
}
