package handlers

import (
	"encoding/json"
	"github.com/shirou/gopsutil/host"
	"log"
	"net"
	"os"
	"runtime"
	"syscall"
	"time"
)

//DiskStatus holds the disk information
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

//SysInfo holds the system info where in the that the app is running
type SysInfo struct {
	CurrentUTC          time.Time `json:"currentUTC"`
	CurrentLocalTime    time.Time `json:"currentLocalTime"`
	GolangVersion       string    `json:"golangVersion"`
	ContainerHostName   string    `json:"containerHostName"`
	Uptime              uint64    `json:"upTime"`
	UptimeDays          uint64    `json:"uptimeDays"`
	UptimeHours         uint64    `json:"uptimeHours"`
	UptimeMinutes       uint64    `json:"uptimeMinutes"`
	CPUs                int       `json:"CPUs"`
	AllocMemory         uint64    `json:"allocatedMemory"`
	AllocMemoryMB       uint64    `json:"allocatedMemoryMB"`
	TotalAllocMemory    uint64    `json:"totalAllocatedMemory"`
	TotalAllocMemoryMB  uint64    `json:"totalAllocatedMemoryMB"`
	TotalSystemMemory   uint64    `json:"totalSystemMem"`
	TotalSystemMemoryMB uint64    `json:"totalSystemMemMB"`
	NetworkInterfaces   [20]struct {
		Name            string `json:"networkInterfaceName"`
		HardwareAddress string `json:"hardwareAddress"`
		IPAddresses     [5]struct {
			IPAddress string `json:"ipAddress"`
		} `json:"ipAddresses"`
	} `json:"networkInterfaces"`
	Disk [10]struct {
		Path string  `json:"diskPath"`
		All  float64 `json:"TotalStorage"`
		Used float64 `json:"UsedStorage"`
		Free float64 `json:"FreeStorage"`
	}
}

//SystemInfo will return various information about the sytem (VM or container) in which it is running
func SystemInfo() string {

	var si SysInfo
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	si.AllocMemory = m.Alloc
	si.AllocMemoryMB = btomb(m.Alloc)
	si.TotalAllocMemory = m.TotalAlloc
	si.TotalAllocMemoryMB = btomb(m.TotalAlloc)
	si.TotalSystemMemory = m.Sys
	si.TotalSystemMemoryMB = btomb(m.Sys)
	si.CPUs = runtime.NumCPU()

	si.GolangVersion = runtime.Version()
	si.ContainerHostName, _ = os.Hostname()
	si.CurrentUTC = time.Now().UTC()

	si.CurrentLocalTime = time.Now().Local()

	const (
		B  = 1
		KB = 1024 * B
		MB = 1024 * KB
		GB = 1024 * MB
	)

	si.Uptime, _ = host.Uptime()
	si.UptimeDays = si.Uptime / (60 * 60 * 24)
	si.UptimeHours = (si.Uptime - (si.UptimeDays * 60 * 60 * 24)) / (60 * 60)
	si.UptimeMinutes = ((si.Uptime - (si.UptimeDays * 60 * 60 * 24)) - (si.UptimeHours * 60 * 60)) / 60
	interfaces, err := net.Interfaces()

	if err == nil {
		for i, interfac := range interfaces {
			if interfac.Name == "" {
				continue
			}
			addrs, _ := interfac.Addrs()
			si.NetworkInterfaces[i].Name = interfac.Name
			si.NetworkInterfaces[i].HardwareAddress = string(interfac.HardwareAddr)
			for x, addr := range addrs {
				if addr.String() != "" {
					si.NetworkInterfaces[i].IPAddresses[x].IPAddress = addr.String()
				} else {
					break
				}
			}
		}
	}

	var paths [10]string
	paths[0] = "/"

	for i, path := range paths {
		disk := DiskUsage(path)
		si.Disk[i].Path = path
		si.Disk[i].All = float64(disk.All) / float64(GB)
		si.Disk[i].Used = float64(disk.Used) / float64(GB)
		si.Disk[i].Free = float64(disk.Free) / float64(GB)
	}

	strJSON, err := json.Marshal(si)
	checkErr(err)

	return string(strJSON)
}

func btomb(b uint64) uint64 {
	return b / 1024
}

// DiskUsage disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {

	if path == "" {
		return
	}

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		log.Fatal(err)
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return disk
}
