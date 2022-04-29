package memory

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetMemory(t *testing.T) {
	res, err := Get()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	require.Nil(t, err)
}
