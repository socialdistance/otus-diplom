package memory

import (
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

type Stats struct {
	Total, Used, Free float64
}

// Get TODO: done this
func Get() (*Stats, error) {
	ret, err := unix.SysctlRaw("vm.swapusage")
	if err != nil {
		return nil, fmt.Errorf("failed in sysctl vm.swapusage: %s", err)
	}

	load := *(*Stats)(unsafe.Pointer(&ret[0]))

	// get percent (load.Used * 100) / load.Total
	return &Stats{
		Total: load.Total,
		Used:  (load.Used * 100) / load.Total,
		Free:  load.Free,
	}, nil
}
