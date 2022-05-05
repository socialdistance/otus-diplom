package average

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Run("loadavg positive test", func(t *testing.T) {
		res, err := Get()
		if err != nil {
			t.Fatal(err)
		}

		require.NoError(t, err)
		require.NotNil(t, res.Loadavg1)
		require.IsType(t, 1.0, res.Loadavg1)
		require.NotNil(t, res.Loadavg5)
		require.IsType(t, 1.0, res.Loadavg5)
		require.NotNil(t, res.Loadavg15)
		require.IsType(t, 1.0, res.Loadavg15)
	})
}
