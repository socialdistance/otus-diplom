package cpu

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// #include <mach/mach_host.h>
// #include <mach/host_info.h>
import "C"

func Get() (*Stats, error) {
	return collectCPUStats()
}

type Stats struct {
	User, System, Idle int
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
		User:   user,
		System: system,
		Idle:   idle,
	}, nil
}
