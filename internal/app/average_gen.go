package app

import (
	average "static_collector/internal/gather/average"
)

type loadavgGenerator struct{}

func (gen *loadavgGenerator) Get() (metric, error) {
	loadavg, err := average.Get()

	return metric{
		Name: "loadavg",
		values: []Value{
			{"loadavg.1m", loadavg.Loadavg1, "-"},
			{"loadavg.5m", loadavg.Loadavg5, "-"},
			{"loadavg.15m", loadavg.Loadavg15, "-"}},
	}, err
}
