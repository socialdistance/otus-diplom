package network

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Stats struct {
	Protocol string
	RecvQ    string
	SendQ    string
	Local    string
	Foreign  string
	State    string
}

func Get() ([]Stats, error) {
	// netstat -anv -p tcp -> -p tcp dont work ?
	cmd := exec.Command("netstat", "-anv")
	var out io.Reader
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	out = io.MultiReader(stdout, stderr)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	network, err := collectNetworkStats(out)
	if err != nil {
		return nil, err
	}

	return network, nil
}

func collectNetworkStats(out io.Reader) ([]Stats, error) {
	scanner := bufio.NewScanner(out)
	var network []Stats
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		protocol := strings.Fields(fields[0])
		if strings.HasPrefix(protocol[0], "tcp4") || strings.HasPrefix(protocol[0], "udp4") {
			network = append(network, Stats{
				Protocol: fields[0],
				RecvQ:    fields[1],
				SendQ:    fields[2],
				Local:    fields[3],
				Foreign:  fields[4],
				State:    fields[5],
			})
		}
	}

	return network, nil
}
