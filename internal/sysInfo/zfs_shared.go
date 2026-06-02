package sysinfo

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func ZFSDetect() bool {
	if runtime.GOOS == "linux" {
		content := ReadFile("/proc/modules")
		return strings.Contains(content, "zfs")
	}
	return runtime.GOOS == "freebsd"
}

func ZFSPoolInfo() string {
	out := ExecCommand("zpool", "list", "-H", "-o", "name,size,alloc,capacity")
	fields := strings.Fields(out)
	if len(fields) < 4 {
		return "Unknown"
	}
	return fmt.Sprintf("%s / %s (used: %s)", fields[2], fields[1], fields[3])
}

func ZFSArcSize() uint64 {
	if runtime.GOOS == "linux" {
		content := ReadFile("/proc/spl/kstat/zfs/arcstats")
		for line := range strings.SplitSeq(content, "\n") {
			if strings.HasPrefix(line, "size ") {
				fields := strings.Fields(line)
				if len(fields) >= 3 {
					val, _ := strconv.ParseUint(fields[2], 10, 64)
					return val
				}
			}
		}
	}
	if runtime.GOOS == "freebsd" {
		out := ExecCommand("sysctl", "-n", "kstat.zfs.misc.arcstats.size")
		val, _ := strconv.ParseUint(strings.TrimSpace(out), 10, 64)
		return val
	}
	return 0
}
