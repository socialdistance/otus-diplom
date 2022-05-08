//go:build darwin
// +build darwin

package disk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDisk(t *testing.T) {
	t.Run("disk test", func(t *testing.T) {
		res, err := Get()
		require.Nil(t, err)

		if res.Tps <= 0 || res.Mb <= 0 || res.Kb <= 0 {
			t.Errorf("invalid disk value: %+v", res)
		}
	})
}
