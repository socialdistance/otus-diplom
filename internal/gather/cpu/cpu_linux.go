//go:build linux
// +build linux

package cpu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type StatsLinux struct {
	User, System, Idle float64
}

type cpuStat struct {
	name string
	ptr  *float64
}

func Get() (*StatsLinux, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return collectCPULinuxStats(file)
}

func collectCPULinuxStats(out io.Reader) (*StatsLinux, error) {
	scanner := bufio.NewScanner(out)
	var cpu StatsLinux

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to scan /proc/stat")
	}

	cpuStats := []cpuStat{
		{"user", &cpu.User},
		{"system", &cpu.System},
		{"idle", &cpu.Idle},
	}

	valStrs := strings.Fields(scanner.Text())[1:4]

	for i, valStr := range valStrs {
		val, err := strconv.ParseUint(valStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to scan %s from /proc/stat", cpuStats[i].name)
		}
		*cpuStats[i].ptr = float64(val)
	}

	return &cpu, nil
}
