//go:build unix

package sysinfo

import (
	"fmt"
	"syscall"
)

func (s *SystemInfo) GetDisk() {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		s.Disk = "Unknown"
		return
	}
	totalBytes := stat.Blocks * uint64(stat.Bsize)
	usedBytes := totalBytes - (stat.Bavail * uint64(stat.Bsize))
	const bToGb = 1073741824.0
	diskTotal := float64(totalBytes) / bToGb
	diskUsed := float64(usedBytes) / bToGb
	diskPorc := (diskUsed / diskTotal) * 100
	s.Disk = fmt.Sprintf("%.1f / %.1f (used: %.1f%%)", diskUsed, diskTotal, diskPorc)
}
