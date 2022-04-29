package app

import (
	disk "static_collector/internal/gather/disk"
)

type diskGenerator struct {
	disks *disk.Stats
	err   error
}

func (gen *diskGenerator) Get() {
	gen.disks, gen.err = disk.Get()
}

func (gen *diskGenerator) Error() error {
	return gen.err
}

func (gen *diskGenerator) Print(out chan<- value) {
	out <- value{"disk kb", gen.disks.Kb, "-"}
	out <- value{"disk tps", gen.disks.Tps, "-"}
	out <- value{"disk mb", gen.disks.Mb, "-"}
}
