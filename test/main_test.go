package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	var fortyTwo = 42
	if (fortyTwo != 42) {
		t.Error("fail")
	}
}