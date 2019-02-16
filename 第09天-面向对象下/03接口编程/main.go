package main

type inter interface {
	socket()
	ency()
}

func framework(i inter) {
	i.socket()
	i.ency()
}
func main() {
	f1 := factory1{"工厂1"}
	f2 := factory2{"工厂2"}
	framework(&f1)
	framework(&f2)
}
