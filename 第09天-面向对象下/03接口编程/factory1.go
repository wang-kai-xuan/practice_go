package main

import "fmt"

type factory1 struct {
	d string
}

func (f *factory1) socket() {
	fmt.Println("socket:", *f)
}

func (f *factory1) ency() {
	fmt.Println("ency:", *f)
}
