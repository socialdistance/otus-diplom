package app

import (
	"static_collector/internal/config"
)

var generators []generator

func InitGenerator(config config.StatsConfig) []generator {
	if config.LoadAvg {
		generators = append(generators, &loadavgGenerator{})
	}

	if config.CPU {
		generators = append(generators, &cpuGenerator{})
	}

	if config.Disk {
		generators = append(generators, &diskGenerator{})
	}

	if config.Memory {
		generators = append(generators, &memoryGenerator{})
	}
	// generators = []generator{
	// &loadavgGenerator{},
	// &cpuGenerator{},
	// &diskGenerator{},
	// &memoryGenerator{},
	// &networkGenerator{},
	// &talkersGenerator{},
	// }

	return generators
}
