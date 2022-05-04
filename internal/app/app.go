package app

import (
	"context"
	"fmt"
	"log"
	"static_collector/internal/config" // nolint:gci
	"sync"
	"time"
)

type generator interface {
	Get() (metric, error)
}

type Value struct {
	Name  string
	Value interface{}
	unit  string //nolint:structcheck
}

type metric struct {
	Name   string
	values []Value
}

func Run(ctx context.Context, n, m int64, config config.Config) chan map[string][][]Value {
	var mutex sync.Mutex

	mapSlice := gatherGenerators(ctx, &mutex, config)

	result := gatherResult(ctx, mapSlice, n, m, &mutex)

	return result
}

func CalculateRes(keyCount int, value [][]Value, key string, m int64) []string {
	var tmpResult []string
	for i := 0; i < keyCount; i++ {
		metricAverage := 0.
		for j := 0; j < len(value); j++ {
			metricAverage += (value[j][i].Value).(float64)
		}
		tmpString := fmt.Sprintf("metric: %s, kind: %s - %f", key, value[0][i].Name, metricAverage/float64(m))
		tmpResult = append(tmpResult, tmpString)
	}

	return tmpResult
}

func gatherResult(ctx context.Context, mapSlice map[string][]metric, n, m int64, mutex *sync.Mutex) chan map[string][][]Value { //nolint:lll
	result := make(chan map[string][][]Value)

	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(n))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				tmpMap := make(map[string][][]Value)
				mutex.Lock()
				for name, metrics := range mapSlice {
					if int64(len(metrics)) > m {
						values := make([][]Value, 0, m)
						for _, item := range metrics[:m] {
							values = append(values, item.values)
						}
						tmpMap[name] = values
					}
				}
				mutex.Unlock()
				result <- tmpMap
			case <-ctx.Done():
				close(result)
			}
		}
	}()

	return result
}

func gatherGenerators(ctx context.Context, mutex *sync.Mutex, config config.Config) map[string][]metric {
	mapSlice := make(map[string][]metric)

	generators := InitGenerator(config) //nolint:typecheck

	for _, gen := range generators {
		ticker := time.NewTicker(time.Second * 1)
		go func(gen generator, ticker *time.Ticker) {
			for {
				select {
				case <-ticker.C:
					value, err := gen.Get()
					if err != nil {
						log.Printf("error get from generator %s", err)
						continue
					}
					mutex.Lock()
					if _, ok := mapSlice[value.Name]; !ok {
						mapSlice[value.Name] = make([]metric, 0)
					}
					mapSlice[value.Name] = append(mapSlice[value.Name], value)
					mutex.Unlock()
				case <-ctx.Done():
					log.Printf("Client abort")
					ticker.Stop()
					return
				}
			}
		}(gen, ticker)
	}

	return mapSlice
}
