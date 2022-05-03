package app

import (
	"static_collector/internal/gather/cpu"
)

type cpuGenerator struct{}

func (gen *cpuGenerator) Get() (metric, error) {
	cpu, err := cpu.Get()

	return metric{
		Name: "cpu",
		values: []Value{
			{"cpu.user", cpu.User, "-"},
			{"cpu.system", cpu.System, "-"},
			{"cpu.idle", cpu.Idle, "-"},
		},
	}, err
}
