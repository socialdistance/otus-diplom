package disk

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Stats struct {
	Kb  float64
	Tps int
	Mb  float64
}

func Get() (*Stats, error) {
	cmd, err := exec.Command("iostat").Output()
	if err != nil {
		log.Fatal(err)
	}

	res := strings.Fields(string(cmd))

	kb, err := strconv.ParseFloat(res[13], 64)
	if err != nil {
		return nil, err
	}

	tp, err := strconv.Atoi(res[14])
	if err != nil {
		return nil, err
	}

	mb, err := strconv.ParseFloat(res[15], 64)
	if err != nil {
		return nil, err
	}

	return collectDiskUsage(kb, tp, mb)
}

func collectDiskUsage(kb float64, tp int, mb float64) (*Stats, error) {
	return &Stats{
		Kb:  kb,
		Tps: tp,
		Mb:  mb,
	}, nil
}
