package main

import (
	"fmt"
	"time"
)

func main() {
	mt := time.NewTimer(time.Second * 10)
	mt.Reset(1 * time.Second)
	go func() {
		res := <-mt.C
		fmt.Println("--child---,time over,res=", res)
	}()
	for {

	}
}
