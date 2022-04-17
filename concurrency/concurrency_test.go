//go:build !go1.18

package concurrency

import (
	"testing"
)

func TestMapSimple(t *testing.T) {
	items := []interface{}{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	results := Map(items, NoOp, 5)
	if len(results) != len(items) {
		t.Errorf("expected results length to be %v, actual %v", len(items), len(results))
	}
}

func BenchmarkMapSimple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		items := []interface{}{1}
		Map(items, NoOp, 5)
	}
}

func NoOp(i interface{}) (interface{}, error) {
	return i, nil
}
