package app

import "static_collector/internal/gather/memory"

type memoryGenerator struct{}

func (gen *memoryGenerator) Get() (metric, error) {
	memoryInfo, err := memory.Get()

	return metric{
		Name: "memory",
		values: []Value{
			{"memory total", memoryInfo.Total, "-"},
			{"memory used %", memoryInfo.Used, "-"},
			{"memory free", memoryInfo.Free, "-"},
		},
	}, err
}
