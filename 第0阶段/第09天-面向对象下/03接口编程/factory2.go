package main

import "fmt"

type factory2 struct {
	d string
}

func (f *factory2) socket() {
	fmt.Println("socket:", *f)
}

func (f *factory2) ency() {
	fmt.Println("ency:", *f)
}
