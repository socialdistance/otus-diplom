package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	res, err := LoadConfig("./configs/config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	require.Nil(t, err)
}
