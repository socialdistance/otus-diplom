//go:build linux
// +build linux

package network

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNetworks(t *testing.T) {
	res, err := Get()
	require.NoError(t, res)
	require.Nil(t, err)
}
