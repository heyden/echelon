package main

import (
	"fmt"
	"time"

	"github.com/heyden/echelon/concurrency"
)

func main() {
	items := []interface{}{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	results := concurrency.Map(items, PrintItem, 5)
	fmt.Println(results)
}

// PrintItem prints an item.
func PrintItem(i interface{}) (interface{}, error) {
	time.Sleep(time.Second * 1)
	fmt.Printf("hey it is %v\n", i)
	return i, nil
}
