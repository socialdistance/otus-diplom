package talkers

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTalkers(t *testing.T) {
	res, err := Get()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	require.Nil(t, err)
}

func TestGetTalkersBytes(t *testing.T) {
	res, err := GetByte()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	require.Nil(t, err)
}

func TestGetTalkersResult(t *testing.T) {
	res, err := Result()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	require.Nil(t, err)
}
