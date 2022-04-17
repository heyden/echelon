//go:build !go1.18

package concurrency

import (
	"fmt"
	"sync"
)

type MapperFunction func(item interface{}) (interface{}, error)

func Map(items []interface{}, fn MapperFunction, concurrency int) []interface{} {
	done := make(chan bool, 1)
	results := make(chan interface{}, len(items))
	defer close(results)
	defer close(done)

	mapper := func(items []interface{}, fn MapperFunction, results chan interface{}) {
		semaphore := make(chan struct{}, concurrency)
		defer close(semaphore)
		wg := sync.WaitGroup{}

		for _, v := range items {
			wg.Add(1)
			semaphore <- struct{}{}

			go func(fn MapperFunction, value interface{}, semaphore chan struct{}) {
				defer func() {
					<-semaphore
					wg.Done()
				}()

				r, err := fn(value)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
				results <- r
			}(fn, v, semaphore)
		}
		wg.Wait()
		done <- true
	}

	consumer := func(done chan bool, results chan interface{}) []interface{} {
		data := []interface{}{}
		for {
			select {
			case r := <-results:
				data = append(data, r)
			case <-done:
				return data
			}
		}
	}

	go mapper(items, fn, results)
	data := consumer(done, results)
	return data
}
