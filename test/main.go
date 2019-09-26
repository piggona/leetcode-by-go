package main

import "fmt"

type f2 interface {
	F2()
}

type A struct{}

func (a A) F1(f f2) {
	fmt.Println("A.F1")
	f.F2()
}

func (a A) F2() {
	fmt.Println("A.F2")
}

type B struct {
	A
}

func (b B) F2() {
	fmt.Println("B.F2")
}

func main() {
	d := B{}
	// value := d.(A)
	d.F1(d)
	d.F1(d.A)
}
