package average

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Stats struct {
	Loadavg1, Loadavg5, Loadavg15 float64
}

type loadStruct struct {
	Ldavg  [3]uint32
	Fscale uint64
}

func Get() (*Stats, error) {
	res, err := unix.SysctlRaw("vm.loadavg") //nolint:typecheck
	if err != nil {
		return nil, fmt.Errorf("failed in sysctl vm.loadavg: %w", err)
	}

	return collectLoadavgStats(res)
}

func collectLoadavgStats(out []byte) (*Stats, error) {
	if len(out) != 24 {
		return nil, fmt.Errorf("unexpected output of sysctl vm.loadavg: %v (len: %d)", out, len(out))
	}
	load := *(*loadStruct)(unsafe.Pointer(&out[0]))
	return &Stats{
		Loadavg1:  float64(load.Ldavg[0]) / float64(load.Fscale),
		Loadavg5:  float64(load.Ldavg[1]) / float64(load.Fscale),
		Loadavg15: float64(load.Ldavg[2]) / float64(load.Fscale),
	}, nil
}
