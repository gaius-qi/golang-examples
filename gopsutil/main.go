package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	// Process.
	proc, _ := process.NewProcess(int32(os.Getpid()))
	pc, _ := proc.CPUPercent()
	fmt.Printf("Proc CPU Percent: %v\n", pc)

	pm, _ := proc.MemoryPercent()
	fmt.Printf("Proc MEM Percent: %v\n", pm)

	pn, _ := net.ConnectionsPid("tcp", 1788)
	for _, conn := range pn {
		fmt.Printf("Proc Proc Net: %v\n", conn)
	}

	ov, _ := mem.VirtualMemory()
	fmt.Printf("OS MEM Total: %v, Available:%v, UsedPercent:%f%%\n", ov.Total, ov.Available, ov.UsedPercent)

	oc, _ := cpu.Percent(0, false)
	fmt.Printf("OS CPU Total Percent: %v\n", oc[0])

	oct, _ := cpu.Times(false)
	fmt.Printf("OS CPU Times: %v", oct)

	on, _ := net.Connections("tcp")
	fmt.Printf("OS Proc Net: %v\n", len(on))

	d, _ := disk.Usage("/Users/qiwenbo/.dragonfly/data")
	fmt.Printf("Dist: %v\n", d)

	h, _ := host.Info()
	fmt.Printf("Host: %v\n", h)
}
