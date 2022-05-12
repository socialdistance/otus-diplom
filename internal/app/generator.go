package app

import (
	"static_collector/internal/config"
)

var generators, generatorsLinux []generator

func InitGenerator(config config.Stats) []generator {
	if config.LoadAvg {
		generators = append(generators, &loadavgGenerator{})
	}

	if config.CPU {
		generators = append(generators, &cpuGenerator{})
	}

	if config.Disk {
		generators = append(generators, &diskGenerator{})
	}

	return generators
}

func InitGeneratorLinux(config config.Stats) []generator {
	if config.LoadAvg {
		generatorsLinux = append(generatorsLinux, &loadavgLinuxGenerator{})
	}

	if config.CPU {
		generatorsLinux = append(generatorsLinux, &cpuLinuxGenerator{})
	}

	if config.NetStat {
		generatorsLinux = append(generatorsLinux, &networkLinuxGenerator{})
	}

	return generatorsLinux
}
