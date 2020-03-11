package monitor


import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)


func Hello() string{
	return "Hello"
}

func WatchHost()(*host.InfoStat, error){
	return host.Info()
}

func WatchCpu()([]cpu.InfoStat,error){
	return cpu.Info()
}

func WatchDisk(path string)  (*disk.UsageStat, error){
	return disk.Usage(path)
}

func WatchMem() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

func WatchProcess() ([]int32, error) {
	return process.Pids()
}

func WatchNet() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}