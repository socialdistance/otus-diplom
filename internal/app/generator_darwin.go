package app

import (
	"static_collector/internal/config"
)

var generators []generator

func InitGenerator(config config.Config) []generator {
	if config.Stats.LoadAvg {
		generators = append(generators, &loadavgGenerator{})
	}

	if config.Stats.CPU {
		generators = append(generators, &cpuGenerator{})
	}

	if config.Stats.Disk {
		generators = append(generators, &diskGenerator{})
	}

	if config.Stats.Memory {
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
