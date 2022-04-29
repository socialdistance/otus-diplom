package talkers

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Network struct {
	Name        string
	Source      string
	Destination string
	Bps         string
}

type StatsBytes struct {
	Name             string
	RxBytes, TxBytes uint64
}

type Stats struct {
	Network    []Network
	StatsBytes []StatsBytes
}

func Result() (*Stats, error) {
	network, err := Get()
	if err != nil {
		return nil, err
	}

	statsBytes, err := GetByte()
	if err != nil {
		return nil, err
	}

	return &Stats{
		Network:    network,
		StatsBytes: statsBytes,
	}, nil
}

func Get() ([]Network, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "nettop", "-l1 -m route")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	networks, err := collectNetworkStats(out)
	if err != nil {
		return nil, err
	}

	return networks, nil
}

func GetByte() ([]StatsBytes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "netstat", "-bni")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	networks, err := collectNetworkStatsBytes(out)
	if err != nil {
		go cmd.Wait() //nolint:errcheck
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return networks, nil

}

func collectNetworkStatsBytes(out io.Reader) ([]StatsBytes, error) {
	var rxBytesIdx, txBytesIdx int
	var networkBytes []StatsBytes

	scanner := bufio.NewScanner(out)

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to scan output of netstat")
	}

	line := scanner.Text()
	if !strings.HasPrefix(line, "Name") {
		return nil, fmt.Errorf("unexpected output of netstat -bni: %s", line)
	}

	fields := strings.Fields(line)
	fieldsCount := len(fields)

	for i, field := range fields {
		switch field {
		case "Ibytes":
			rxBytesIdx = i
		case "Obytes":
			txBytesIdx = i
		}
	}

	if rxBytesIdx == 0 || txBytesIdx == 0 {
		return nil, fmt.Errorf("unexpected output of netstat -bni: %s", line)
	}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		name := strings.TrimSuffix(fields[0], "*")
		if strings.HasPrefix(name, "lo") || !strings.HasPrefix(fields[2], "<Link#") {
			continue
		}
		rxBytesIdx, txBytesIdx := rxBytesIdx, txBytesIdx
		if len(fields) < fieldsCount { // Address can be empty
			rxBytesIdx, txBytesIdx = rxBytesIdx-1, txBytesIdx-1
		}
		rxBytes, err := strconv.ParseUint(fields[rxBytesIdx], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Ibytes of %s", name)
		}
		txBytes, err := strconv.ParseUint(fields[txBytesIdx], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Obytes of %s", name)
		}
		networkBytes = append(networkBytes, StatsBytes{Name: name, RxBytes: rxBytes, TxBytes: txBytes})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for netstat: %s", err)
	}

	return networkBytes, nil

}

func collectNetworkStats(out io.Reader) ([]Network, error) {
	var networks []Network

	scanner := bufio.NewScanner(out)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		name := strings.TrimSuffix(fields[1], "*")
		if strings.HasPrefix(name, "tcp4") || strings.HasPrefix(name, "udp4") {
			if len(fields) > 4 {
				bytes, err := strconv.ParseUint(fields[3], 10, 64)
				if err != nil {
					continue
				}

				networks = append(networks, Network{
					Name:        name,
					Source:      fields[2],
					Destination: fields[3],
					Bps:         fmt.Sprintf("%d %s", bytes, fields[4]),
				})
			}
		}
	}

	sort.Slice(networks, func(i, j int) bool {
		return networks[i].Bps > networks[j].Bps
	})

	return networks, nil
}
