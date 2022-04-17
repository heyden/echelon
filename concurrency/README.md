# Map with rate limit

Map capability with rate limit similar to Bluebird's `Promise.map`.

## Go Playground

With `interface{}` - https://go.dev/play/p/zQGmKm_BjCv

With go1.18 generics - https://go2goplay.golang.org/p/TodjZGwrZ4Q

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/heyden/go-concurrency/concurrency"
)

func main() {
	items := []interface{}{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	results := concurrency.Map(items, PrintItem, 5)
	fmt.Println(results)
}

func PrintItem(i interface{}) (interface{}, error) {
	time.Sleep(time.Second * 1)
	fmt.Printf("hey it is %v\n", i)
	return i, nil
}
```