package app

import (
	"static_collector/internal/gather/talkers"
)

type talkersGenerator struct {
	talker *talkers.Stats
	error  error
}

func (gen *talkersGenerator) Get() {
	gen.talker, gen.error = talkers.Result()
}

func (gen *talkersGenerator) Error() error {
	return gen.error
}

func (gen *talkersGenerator) Print(out chan<- value) {
	for _, statsByte := range gen.talker.StatsBytes {
		out <- value{"network." + statsByte.Name + ".rx_bytes", statsByte.RxBytes, "bytes"}
		out <- value{"network." + statsByte.Name + ".tx_bytes", statsByte.TxBytes, "bytes"}
	}

	for _, network := range gen.talker.Network {
		out <- value{"network." + network.Name + " " + network.Source + " " + network.Destination, network.Bps, "bytes"}
	}
}
