package main

import (
	"context"
	"fmt"
	"io"
	"os"

	httpx "github.com/heyden/echelon/http"
)

func main() {
	resp, err := httpx.Get(context.TODO(), "https://google.com", nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
