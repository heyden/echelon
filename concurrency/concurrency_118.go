//go:build go1.18

package concurrency

import (
	"fmt"
	"sync"
)

type MapperFunction[T, R any] func(item T) (R, error)

func Map[T, R any](items []T, fn MapperFunction[T, R], concurrency int) []R {
	done := make(chan bool, 1)
	results := make(chan R)
	defer close(results)
	defer close(done)

	mapper := func(items []T, fn MapperFunction[T, R], results chan R) {
		semaphore := make(chan struct{}, concurrency)
		defer close(semaphore)
		wg := sync.WaitGroup{}

		for _, v := range items {
			wg.Add(1)
			semaphore <- struct{}{}

			go func(fn MapperFunction[T, R], value T, semaphore chan struct{}) {
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

	consumer := func(done chan bool, results chan R) []R {
		data := []R{}
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
