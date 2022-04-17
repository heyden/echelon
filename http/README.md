# Simple HTTP package that wraps `net/http`

## Usage

See the [examples](./examples) directory for more examples.

```go
package main

import (
	"context"
	"fmt"
	"io"
	"os"

	httpx "github.com/heyden/go-http/http"
)

func main() {
	resp, err := httpx.Get(context.TODO(), "https://google.com", nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}

```