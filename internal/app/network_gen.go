package app

import "static_collector/internal/gather/network"

type networkLinuxGenerator struct{}

func (gen *networkLinuxGenerator) Get() (metric, error) {
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
