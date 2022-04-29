package app

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	out := bufio.NewWriter(new(bytes.Buffer))
	errs := Run(nil, out)
	if errs != nil {
		t.Errorf("error occured: %v", errs)
	}
	fmt.Println(errs)
}
