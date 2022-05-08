//go:build linux
// +build linux

package average

import (
	"fmt"
	"io"
	"os"
)

type StatsLinux struct {
	Loadavg1, Loadavg5, Loadavg15 float64
}

func Get() (*StatsLinux, error) {
	file, err := os.Open("/proc/loadavg")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collectLoadavgLinuxStats(file)
}

func collectLoadavgLinuxStats(out io.Reader) (*StatsLinux, error) {
	var loadavg StatsLinux
	ret, err := fmt.Fscanf(out, "%f %f %f", &loadavg.Loadavg1, &loadavg.Loadavg5, &loadavg.Loadavg15)
	if err != nil || ret != 3 {
		return nil, fmt.Errorf("unexpected format of /proc/loadavg")
	}
	return &loadavg, nil
}
