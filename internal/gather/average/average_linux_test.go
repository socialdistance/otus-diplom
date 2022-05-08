//go:build linux
// +build linux

package average

import (
	"testing"
)

func TestGetLoadavg(t *testing.T) {
	loadavg, err := GetLinux()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	if loadavg.Loadavg1 < 0 || loadavg.Loadavg5 < 0 || loadavg.Loadavg15 < 0 {
		t.Errorf("invalid loadavg value: %v", loadavg)
	}
	t.Logf("loadavg value: %+v", loadavg)
}