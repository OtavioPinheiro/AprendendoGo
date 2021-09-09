//Links: https://godoc.org/testing
// https://www.golang-book.com/books/intro/12

package main

import (
	"testing"
)

func TestMedia(t *testing.T) {
	var v float64
	v = Media([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
