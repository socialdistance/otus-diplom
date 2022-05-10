package network

import (
	"os/exec"
	"strconv"
	"strings"
)

type Stats struct {
	Protocol string
	RecvQ    float64
	SendQ    float64
	Local    string
	Foreign  string
	State    string
}

func Get() (*Stats, error) {
	cmd, err := exec.Command("netstat", "-lntup").Output()
	if err != nil {
		return nil, err
	}

	res := strings.Fields(string(cmd))

	recvq, err := strconv.ParseFloat(res[16], 64)
	if err != nil {
		return nil, err
	}

	sendq, err := strconv.ParseFloat(res[17], 64)
	if err != nil {
		return nil, err
	}

	return &Stats{
		Protocol: res[15],
		RecvQ:    recvq,
		SendQ:    sendq,
		Local:    res[18],
		Foreign:  res[19],
		State:    res[20],
	}, nil
}
