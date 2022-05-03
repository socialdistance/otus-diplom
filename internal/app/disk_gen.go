package app

import (
	disk "static_collector/internal/gather/disk"
)

type diskGenerator struct{}

func (gen *diskGenerator) Get() (metric, error) {
	disks, err := disk.Get()

	return metric{
		Name: "disk",
		values: []Value{
			{"disk kb", disks.Kb, "-"},
			{"disk tps", disks.Tps, "-"},
			{"disk mb", disks.Mb, "-"},
		},
	}, err
}
