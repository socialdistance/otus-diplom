package config

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoad(t *testing.T) {
	res, err := LoadConfig("./configs/config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	require.Nil(t, err)
}
