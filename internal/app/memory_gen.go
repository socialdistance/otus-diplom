package app

import "static_collector/internal/gather/memory"

type memoryGenerator struct {
	memoryInfo *memory.Stats
	err        error
}

func (gen *memoryGenerator) Get() {
	gen.memoryInfo, gen.err = memory.Get()
}

func (gen *memoryGenerator) Error() error {
	return gen.err
}

func (gen *memoryGenerator) Print(out chan<- value) {
	memory := gen.memoryInfo

	out <- value{"memory total", memory.Total, "-"}
	out <- value{"memory used %", memory.Used, "-"}
	out <- value{"memory free", memory.Free, "-"}
}
