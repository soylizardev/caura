package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func (i *InfoPc) GetPc() {
	vendorRaw, _ := os.ReadFile("/sys/class/dmi/id/sys_vendor")
	nameRaw, _ := os.ReadFile("/sys/class/dmi/id/product_version")
	modelRaw, _ := os.ReadFile("/sys/class/dmi/id/product_name")
	vendor, name, model := strings.TrimSpace(string(vendorRaw)), strings.TrimSpace(string(nameRaw)), strings.TrimSpace(string(modelRaw))
	i.Pc = fmt.Sprintf("%s %s (%s)", vendor, name, model)
}

func (i *InfoPc) GetCPU() {
	cmd, _ := os.ReadFile("/proc/cpuinfo")
	split := strings.Split(string(cmd), "\n")
	var cpuModel string
	var cpuCores string
	cpuThreads := 0

	for _, line := range split {
		if strings.Contains(line, "model name") {
			div := strings.Split(line, ":")
			cpuModel = strings.TrimSpace(div[1])
			cpuThreads++
		}
		if strings.Contains(line, "cpu cores") {
			div := strings.Split(line, ":")
			cpuCores = strings.TrimSpace(div[1])
		}
	}
	i.CPU = fmt.Sprintf("%s (%s Cores / %d Threads)", cpuModel, cpuCores, cpuThreads)
}

func (i *InfoPc) GetGraphics() {
	cmd := exec.Command("sh", "-c", "lspci | grep -i vga")
	out, _ := cmd.Output()
	split := strings.Split(string(out), ":")
	i.Graphic = strings.TrimSpace(split[2])
}

func (i *InfoPc) GetDisk() {
	cmd := exec.Command("df", "-h", "/")
	com, _ := cmd.Output()
	lines := strings.Split(string(com), "\n")
	split := strings.Fields(lines[1])
	diskTotal := split[1]
	diskUsed := split[2]
	diskporc := split[4]

	i.Disk = fmt.Sprintf("%s / %s (used: %s)", diskUsed, diskTotal, diskporc)
}

func (i *InfoPc) GetRam() {
	content, _ := os.ReadFile("/proc/meminfo")
	lines := strings.Split(string(content), "\n")
	var memTotalRaw, memAvailRaw string
	for _, line := range lines {
		if strings.Contains(line, "MemTotal:") {
			fields := strings.Fields(line) // Separa automáticamente por cualquier cantidad de espacios
			memTotalRaw = fields[1]        // El índice 1 contiene el número limpio
		}
		if strings.Contains(line, "MemAvailable:") {
			fields := strings.Fields(line)
			memAvailRaw = fields[1]
		}
	}
	totalMemory, _ := strconv.ParseFloat(memTotalRaw, 64)
	totalGB := totalMemory / 1048576
	availableMemory, _ := strconv.ParseFloat(memAvailRaw, 64)
	availableGB := availableMemory / 1048576
	memUsed := totalGB - availableGB

	i.Ram = fmt.Sprintf("%.2f GiB / %.2f GiB", memUsed, totalGB)
}

func (i *InfoPc) GetSwap() {
	content, _ := os.ReadFile("/proc/meminfo")
	lines := strings.Split(string(content), "\n")
	var swapTotalRaw, swapFreeRaw string
	for _, line := range lines {
		if strings.Contains(line, "SwapTotal:") {
			fields := strings.Fields(line)
			swapTotalRaw = fields[1]
		}
		if strings.Contains(line, "SwapFree:") {
			fields := strings.Fields(line)
			swapFreeRaw = fields[1]
		}
	}

	totalSwap, _ := strconv.ParseFloat(swapTotalRaw, 64)
	totalGB := totalSwap / 1048576
	freeSwap, _ := strconv.ParseFloat(swapFreeRaw, 64)
	freeGB := freeSwap / 1048576
	swapUsed := totalGB - freeGB
	i.Swap = fmt.Sprintf("%.2f GiB / %.2f GiB", swapUsed, totalGB)
}
