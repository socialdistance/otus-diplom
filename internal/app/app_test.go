package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppRun(t *testing.T) {
	value := [][]Value{
		{
			{"cpu.user", float64(9877), "-"},
			{"cpu.system", float64(361), "-"},
			{"cpu.idle", float64(27323), "-"},
		},
	}

	keyCount := 3
	m := int64(5)

	t.Run("calculate result cpu", func(t *testing.T) {
		res := CalculateRes(keyCount, value, "cpu", m)
		require.True(t, strings.Contains(res[0], "cpu"))
		require.Len(t, res, 3)
		require.NotNil(t, res)
	})

	t.Run("calculate result loadavg", func(t *testing.T) {
		res := CalculateRes(keyCount, value, "loadavg", m)
		require.True(t, strings.Contains(res[0], "loadavg"))
		require.Len(t, res, 3)
		require.NotNil(t, res)
	})
}
