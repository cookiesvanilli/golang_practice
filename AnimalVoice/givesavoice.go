package main

import "fmt"

type Givesavoice interface {
	SayHello()
}

type Cat struct {
    greeting string
}
type Dog struct {
	greeting  string
}
type Bee struct {
	greeting  string
}

func (c Cat) SayHello() {//Cat
	fmt.Println("Cat: MEOW")
}
func (d Dog) SayHello()  {//Dog
	fmt.Println("Dog: WOFF")

}
func (b Bee) SayHello() {//Bird
	fmt.Println("Bee: BZZ")

}
func AnimalVoice(a Givesavoice) {
	a.SayHello()
}

func main() {
	c := Cat{greeting : "meow"}
	AnimalVoice(c)
	d := Dog{greeting : "woff"}
	AnimalVoice(d)
	b := Bee{greeting : "bzz"}
	AnimalVoice(b)
}

