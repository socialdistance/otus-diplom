package memory

import (
	"fmt"
	"unsafe"

	ux "golang.org/x/sys/unix"
)

type Stats struct {
	Total, Used, Free float64
}

type loadStats struct {
	Total, Used, Free uint64
}

func Get() (*Stats, error) {
	ret, err := ux.SysctlRaw("vm.swapusage")
	if err != nil {
		return nil, fmt.Errorf("failed in sysctl vm.swapusage: %w", err)
	}

	load := *(*loadStats)(unsafe.Pointer(&ret[0]))

	// get percent (load.Used * 100) / load.Total
	return &Stats{
		Total: float64(load.Total),
		Used:  float64(load.Used),
		Free:  float64(load.Free),
	}, nil
}
