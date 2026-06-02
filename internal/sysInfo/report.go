package sysinfo

import "fmt"

func (s *SystemInfo) Render() {
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

	//Format
	fmt.Println()
	fmt.Printf("   \033[1;35m%s\033[0m@\033[1;35m%s\033[0m\n", s.User, s.Host)
	fmt.Println("   ------------------OS--------------------")

	// Bloque de Software (OS)
	fmt.Printf("   \033[1;36mOS:\033[0m       %s\n", s.OS)
	fmt.Printf("   \033[1;36mKernel:\033[0m   %s\n", s.Kernel)
	fmt.Printf("   \033[1;36mUptime:\033[0m   %s\n", s.Uptime)
	fmt.Printf("   \033[1;36mShell:\033[0m    %s\n", s.Shell)
	fmt.Printf("   \033[1;36mTerminal:\033[0m %s\n", s.Terminal)
	fmt.Printf("   \033[1;36mIp:\033[0m       %s\n", s.IP)

	fmt.Println("   ------------------PC--------------------")

	// Bloque de Hardware (PC)
	fmt.Printf("   \033[1;33mHost:\033[0m     %s\n", s.Pc)
	fmt.Printf("   \033[1;33mCPU:\033[0m      %s\n", s.CPU)
	fmt.Printf("   \033[1;33mArch:\033[0m     %s\n", s.Arch)
	fmt.Printf("   \033[1;33mGraphics:\033[0m %s\n", s.Graphic)
	fmt.Printf("   \033[1;33mDisk:\033[0m     %s\n", s.Disk)
	fmt.Printf("   \033[1;33mMemory:\033[0m   %s\n", s.Ram)
	fmt.Printf("   \033[1;33mSwap:\033[0m     %s\n", s.Swap)

	fmt.Println("   ----------------------------------------")
	fmt.Println()
}
