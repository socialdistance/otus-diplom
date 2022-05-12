package app

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppRun(t *testing.T) {
	ctx := context.Background()
	var mutex sync.Mutex

	value := [][]Value{
		{
			{"cpu.user", float64(9877), "-"},
			{"cpu.system", float64(361), "-"},
			{"cpu.idle", float64(27323), "-"},
		},
	}

	mapSlice := map[string][]metric{
		"loadavg": {
			{
				Name: "loadavg",
				values: []Value{
					{"cpu.user", float64(9877), "-"},
					{"cpu.system", float64(361), "-"},
					{"cpu.idle", float64(27323), "-"},
				},
			},
		},
	}

	keyCount := 3

	n := int64(5)
	m := int64(5)

	t.Run("test calculate result cpu", func(t *testing.T) {
		res := CalculateRes(keyCount, value, "cpu", m)
		require.True(t, strings.Contains(res[0], "cpu"))
		require.Len(t, res, 3)
		require.NotNil(t, res)
	})

	t.Run("test calculate result loadavg", func(t *testing.T) {
		res := CalculateRes(keyCount, value, "loadavg", m)
		require.True(t, strings.Contains(res[0], "loadavg"))
		require.Len(t, res, 3)
		require.NotNil(t, res)
	})

	t.Run("test gather result", func(t *testing.T) {
		result := gatherResult(ctx, mapSlice, n, m, &mutex)
		require.NotNil(t, result)
	})
}
