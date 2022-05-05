package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Run("invalid config test", func(t *testing.T) {
		_, err := LoadConfig("/tmp/bar.foo")

		require.Error(t, err)
	})
}
