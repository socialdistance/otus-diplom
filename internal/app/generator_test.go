package app

import (
	"static_collector/internal/config" //nolint:gci
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	t.Run("empty test", func(t *testing.T) {
		emptyConfig := config.Stats{
			LoadAvg: false,
			CPU:     false,
			Disk:    false,
			Memory:  false,
			NetStat: false,
			NetTop:  false,
		}

		result := InitGenerator(emptyConfig)
		require.Len(t, result, len(result))
	})

	t.Run("test with config", func(t *testing.T) {
		config := config.Stats{
			LoadAvg: true,
			CPU:     true,
			Disk:    true,
			Memory:  true,
			NetStat: false,
			NetTop:  false,
		}

		result := InitGenerator(config)
		require.Len(t, result, len(result))
	})
}
