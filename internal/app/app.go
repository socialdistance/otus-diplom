package app

import (
	"fmt"
	"sync"
)

type generator interface {
	Get()
	Error() error
	Print(out chan<- value)
}

type value struct {
	name  string
	value interface{}
	unit  string
}

type App struct{}

func (app *App) Run() ([]string, []error) {
	var wg sync.WaitGroup
	var result []string

	for _, gen := range generators {
		wg.Add(1)
		go func(gen generator) {
			defer wg.Done()
			gen.Get()
		}(gen)
	}

	wg.Wait()

	c := make(chan value)
	done := make(chan struct{})
	defer close(done)

	go func() {
		for {
			select {
			case v := <-c:
				result = append(result, fmt.Sprintf("%-25s %-14v %s\n", v.name, v.value, v.unit))
			case <-done:
				close(c)
				return
			}
		}
	}()

	var errs []error
	for _, gen := range generators {
		if err := gen.Error(); err != nil {
			errs = append(errs, err)
		} else {
			gen.Print(c)
		}
	}
	done <- struct{}{}

	return result, errs
}
