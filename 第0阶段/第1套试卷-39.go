package 第0阶段

import "fmt"

type Animal struct {
	species string
}

type Speak interface {
	call()
}

type Cat struct {
	Animal
}
type Dog struct {
	Animal
}

func (c *Dog) call() {
	fmt.Println("汪汪叫")
}
func (c *Cat) call() {
	fmt.Println("喵喵叫")
}

func Call(s Speak) {
	s.call()
}
func main() {
	d := Dog{Animal{"Dog"}}
	c := Cat{Animal{"Cat"}}
	Call(&d)
	Call(&c)
}
