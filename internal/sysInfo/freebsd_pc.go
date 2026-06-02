//go:build freebsd

package sysinfo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s *SystemInfo) GetPc() {
	vendor := ExecCommand("kenv", "-q", "smbios.system.maker")
	product := ExecCommand("kenv", "-q", "smbios.system.product")
	version := ExecCommand("kenv", "-q", "smbios.system.version")

	if vendor == "Unknown" || product == "Unknown" {
		host, _ := os.Hostname()
		model := ExecCommand("sysctl", "-n", "hw.model")
		s.Pc = fmt.Sprintf("%s (%s)", host, model)
		return
	}
	s.Pc = fmt.Sprintf("%s %s (%s)", vendor, product, version)
}

func (s *SystemInfo) GetCPU() {
	cpuModel := ExecCommand("sysctl", "-n", "hw.model")
	cpuThreads := ExecCommand("sysctl", "-n", "hw.ncpu")
	cpuCores := ExecCommand("sysctl", "-n", "kern.sched.topology_spec")
	line := strings.Split(cpuCores, "\n")
	if len(line) < 3 {
		cpuCores = "Unknown"
	}
	for word := range strings.FieldsSeq(line[2]) {
		if strings.HasPrefix(word, `count="`) {
			start := strings.Index(word, `"`)
			end := strings.LastIndex(word, `"`)
			if start != -1 && end != -1 && end > start {
				cpuCores = word[start+1 : end]
			}
			break
		}
	}
	s.CPU = fmt.Sprintf("%s (%s Cores / %s Threads)", cpuModel, cpuCores, cpuThreads)
}

func (s *SystemInfo) GetGraphic() {
	out := ExecCommand("pciconf", "-lv")
	lines := strings.Split(out, "\n")
	inVgaBlock := false
	var vendor, device string

	for _, line := range lines {
		if strings.HasPrefix(line, "vgapci") {
			inVgaBlock = true
			continue
		}
		if inVgaBlock {
			if !strings.HasPrefix(line, "    ") {
				break
			}
			if strings.Contains(line, "vendor") {
				start := strings.Index(line, "'")
				end := strings.LastIndex(line, "'")
				if start != -1 && end != -1 && end > start {
					vendor = line[start+1 : end]
				}
			}
			if strings.Contains(line, "device") {
				start := strings.Index(line, "'")
				end := strings.LastIndex(line, "'")
				if start != -1 && end != -1 && end > start {
					device = line[start+1 : end]
				}
			}
		}
	}

	if vendor == "" && device == "" {
		s.Graphic = "Unknown"
	} else {
		s.Graphic = fmt.Sprintf("%s %s", vendor, device)
	}
}

func (s *SystemInfo) GetRam() {
	totalstr := ExecCommand("sysctl", "-n", "hw.physmem")
	freestr := ExecCommand("sysctl", "-n", "vm.stats.vm.v_free_count")
	inactivestr := ExecCommand("sysctl", "-n", "vm.stats.vm.v_inactive_count")
	pageSizestr := ExecCommand("sysctl", "-n", "hw.pagesize")

	total, errTotal := strconv.ParseFloat(totalstr, 64)
	if errTotal != nil {
		total = 0
	}
	free, errFree := strconv.ParseFloat(freestr, 64)
	if errFree != nil {
		free = 0
	}
	inactive, errInac := strconv.ParseFloat(inactivestr, 64)
	if errInac != nil {
		inactive = 0
	}
	pageSize, errPage := strconv.ParseFloat(pageSizestr, 64)
	if errPage != nil {
		pageSize = 0
	}

	available := (free + inactive) * pageSize
	used := total - available

	const bytesToGiB = 1073741824.0
	totalGiB := total / bytesToGiB
	usedGiB := used / bytesToGiB

	s.Ram = fmt.Sprintf("%.1f GiB / %.1f GiB", usedGiB, totalGiB)
}

func (s *SystemInfo) GetSwap() {
	out := ExecCommand("swapctl", "-l", "-k")
	lines := strings.Split(out, "\n")
	if len(lines) < 2 {
		s.Swap = "Unknown"
		return
	}
	fields := strings.Fields(lines[1])
	if len(fields) < 2 {
		s.Swap = "Unknown"
		return
	}
	totalStr := fields[1]
	usedStr := fields[2]
	totalKB, err := strconv.ParseFloat(totalStr, 64)
	if err != nil {
		s.Swap = "Unknown"
		return
	}
	usedKB, err := strconv.ParseFloat(usedStr, 64)
	if err != nil {
		s.Swap = "Unknown"
		return
	}
	const kbToGiB = 1048576.0
	totalGiB := totalKB / kbToGiB
	usedGiB := usedKB / kbToGiB

	s.Swap = fmt.Sprintf("%.1f GiB / %.1f GiB", usedGiB, totalGiB)
}
