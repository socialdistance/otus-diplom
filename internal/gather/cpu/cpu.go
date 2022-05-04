package cpu

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Stats struct {
	User, System, Idle float64
}

func Get() (*Stats, error) {
	return collectCPUStats()
}

func collectCPUStats() (*Stats, error) {
	cmd, err := exec.Command("iostat").Output()
	if err != nil {
		log.Fatal(err)
	}

	res := strings.Fields(string(cmd))

	user, err := strconv.Atoi(res[16])
	if err != nil {
		return nil, err
	}

	system, err := strconv.Atoi(res[17])
	if err != nil {
		return nil, err
	}

	idle, err := strconv.Atoi(res[18])
	if err != nil {
		return nil, err
	}

	return &Stats{
		User:   float64(user),
		System: float64(system),
		Idle:   float64(idle),
	}, nil
}
