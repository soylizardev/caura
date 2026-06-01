//go:build linux

package sysinfo

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *SystemInfo) GetPc() {
	vendor := ReadFile("/sys/class/dmi/id/sys_vendor")
	name := ReadFile("/sys/class/dmi/id/product_version")
	model := ReadFile("/sys/class/dmi/id/product_name")

	s.Pc = fmt.Sprintf("%s %s (%s)", vendor, name, model)
}

func (s *SystemInfo) GetCPU() {
	cmd := ReadFile("/proc/cpuinfo")
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
	if cpuModel == "" {
		cpuModel = "Unknown"
		cpuThreads = 0
	}
	if cpuCores == "" {
		cpuCores = "0"
	}
	s.CPU = fmt.Sprintf("%s (%s Cores / %d Threads)", cpuModel, cpuCores, cpuThreads)
}

func (s *SystemInfo) GetGraphic() {
	out := ExecCommand("sh", "-c", "lspci | grep -i vga")
	split := strings.Split(out, ":")
	if len(split) > 2 {
		s.Graphic = strings.TrimSpace(string(split[2]))
	} else {
		s.Graphic = "Unknown"
	}
}

func (s *SystemInfo) GetRam() {
	content := ReadFile("/proc/meminfo")

	lines := strings.Split(content, "\n")
	var memTotalRaw, memAvailRaw string
	for _, line := range lines {
		if strings.Contains(line, "MemTotal:") {
			fields := strings.Fields(line)
			memTotalRaw = fields[1]
		}
		if strings.Contains(line, "MemAvailable:") {
			fields := strings.Fields(line)
			memAvailRaw = fields[1]
		}
	}
	totalMemory, errTotal := strconv.ParseFloat(memTotalRaw, 64)
	if errTotal != nil {
		totalMemory = 0
	}
	const kbToGb = 1048576
	totalGB := totalMemory / kbToGb
	availableMemory, errAva := strconv.ParseFloat(memAvailRaw, 64)
	if errAva != nil {
		availableMemory = 0
	}

	availableGB := availableMemory / kbToGb
	memUsed := totalGB - availableGB

	s.Ram = fmt.Sprintf("%.2f GiB / %.2f GiB", memUsed, totalGB)
}

func (s *SystemInfo) GetSwap() {
	content := ReadFile("/proc/meminfo")

	lines := strings.Split(content, "\n")
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

	totalSwap, errTotal := strconv.ParseFloat(swapTotalRaw, 64)
	const kbToGb = 1048576.0
	if errTotal != nil {
		totalSwap = 0
	}
	totalGB := totalSwap / kbToGb
	freeSwap, errFree := strconv.ParseFloat(swapFreeRaw, 64)
	if errFree != nil {
		freeSwap = 0
	}
	freeGB := freeSwap / kbToGb
	swapUsed := totalGB - freeGB
	s.Swap = fmt.Sprintf("%.2f GiB / %.2f GiB", swapUsed, totalGB)
}
