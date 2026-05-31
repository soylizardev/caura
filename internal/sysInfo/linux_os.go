//go:build linux

package sysinfo

func (i *InfoOs) GatherOSInfo() {
	i.GetHost("/proc/sys/kernel/hostname")
	i.GetOS("/etc/os-release", "/proc/sys/kernel/ostype")
	i.GetKernel("/proc/sys/kernel/osrelease")
	i.GetUser()
	i.GetTerm("TERM")
	i.GetShell("SHELL")
	i.GetUptime("/proc/uptime")
}
