package logic

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type CPUInfo struct {
	ModelName string
	Cores     int32
	Speed     float64
}

// SystemStats represents the system stats
type SystemStats struct {
	CPUUsage    []float64
	MemoryUsage float64
	CpuInfo     CPUInfo
}

func GetStats() (SystemStats, error) {
	// get CPU usage
	cpuUsage, err := getCPUUsage()
	if err != nil {
		return SystemStats{}, err
	}

	// get memory usage
	memoryUsage, err := getMemoryUsage()
	if err != nil {
		return SystemStats{}, err
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return SystemStats{}, err
	}

	if len(cpuInfo) == 0 {
		return SystemStats{}, nil
	}

	return SystemStats{
		CPUUsage:    cpuUsage,
		MemoryUsage: memoryUsage,
		CpuInfo: CPUInfo{
			ModelName: cpuInfo[0].ModelName,
			Cores:     cpuInfo[0].Cores,
			Speed:     cpuInfo[0].Mhz / 1000.0, //Convert to GHz
		},
	}, nil

}

func getMemoryUsage() (m float64, err error) {
	// get memory usage
	memory, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}

	return memory.UsedPercent, nil

}

func getCPUUsage() (cpuUsage []float64, err error) {
	// get CPU usage
	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	return cpuPercentages, nil
}
