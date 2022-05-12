//go:build linux
// +build linux

package network

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNetworks(t *testing.T) {
	res, err := Get()

	fmt.Println(res)
	require.Nil(t, err)
}
