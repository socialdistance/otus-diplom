//go:build darwin
// +build darwin

package cpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCPU(t *testing.T) {
	t.Run("cpu test", func(t *testing.T) {
		cpu, err := Get()
		require.Nil(t, err)

		if cpu.User <= 0 || cpu.System <= 0 || cpu.Idle <= 0 {
			t.Errorf("invalid cpu value: %+v", cpu)
		}
	})
}
