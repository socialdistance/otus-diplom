package app

import "static_collector/internal/gather/network"

type networkGenerator struct{}

func (gen *networkGenerator) Get() (metric, error) {
	network, err := network.Get()

	return metric{
		Name: "network",
		values: []Value{
			{"protocol", network.Protocol, "-"},
			{"recvq", network.RecvQ, "-"},
			{"sendq", network.SendQ, "-"},
			{"local", network.Local, "-"},
			{"foreign", network.Foreign, "-"},
			{"state", network.State, "-"},
		},
	}, err
}
