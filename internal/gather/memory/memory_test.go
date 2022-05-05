package memory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMemory(t *testing.T) {
	res, err := Get()
	require.Nil(t, err)

	if res.Total <= 0 || res.Used <= 0 || res.Free <= 0 {
		t.Errorf("invalid disk value: %+v", res)
	}
}
