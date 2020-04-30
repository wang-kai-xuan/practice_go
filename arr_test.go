package main

import (
	"fmt"
	"testing"
)

type OutInfo struct {
	index int
	addr  string
}

func TestArrType(t *testing.T) {
	var outs []OutInfo
	outs = append(outs, OutInfo{1, "item1"})
	outs = append(outs, OutInfo{2, "item2"})

	for _, item := range outs {
		fmt.Println(item)
	}
}
