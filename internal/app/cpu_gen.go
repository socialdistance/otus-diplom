package app

import (
	"static_collector/internal/gather/cpu"
)

type cpuGenerator struct {
	cpu *cpu.Stats
	err error
}

func (gen *cpuGenerator) Get() {
	gen.cpu, gen.err = cpu.Get()
}

func (gen *cpuGenerator) Error() error {
	return gen.err
}

func (gen *cpuGenerator) Print(out chan<- value) {
	cpu := gen.cpu
	out <- value{"cpu.user", cpu.User, "-"}
	out <- value{"cpu.system", cpu.System, "-"}
	out <- value{"cpu.idle", cpu.Idle, "-"}
}
