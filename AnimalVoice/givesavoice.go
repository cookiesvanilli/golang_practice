package main

import "fmt"

type Givesavoice interface {
	SayHello() string
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

func (c *Cat) SayHello() string {//Cat
	fmt.Println("Cat: MEOW")
	return "MEOW"
}
func (d *Dog) SayHello()  string {//Dog
	fmt.Println("Dog: WOFF")
	return "WOFF"

}
func (b *Bee) SayHello() string  {//Bird
	fmt.Println("Bee: BZZ")
	return "BZZ"

}
func AnimalVoice(a Givesavoice) {
	a.SayHello()
}

func main() {
	c := &Cat{greeting : "meow"}
	AnimalVoice(c)
	d := &Dog{greeting : "woff"}
	AnimalVoice(d)
	b := &Bee{greeting : "bzz"}
	AnimalVoice(b)
}

