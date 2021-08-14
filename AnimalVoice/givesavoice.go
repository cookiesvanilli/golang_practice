package main

import "fmt"

type Givesavoice interface {
	sayHello() string
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
func (c *Cat) sayHello() string { //Cat
	fmt.Println("Cat: MEOW")
	return "MEOW"
}
func (d *Dog) sayHello()  string { //Dog
	fmt.Println("Dog: WOFF")
	return "WOFF"
}
func (b *Bee) sayHello() string  { //Bird
	fmt.Println("Bee: BZZ")
	return "BZZ"
}
func animalVoice(list ...Givesavoice) {
	for _, a := range list {
		a.sayHello()
	}
}
func main() {
	c := &Cat{greeting : "meow"}
	d := &Dog{greeting : "woff"}
	b := &Bee{greeting : "bzz"}
	animalVoice(c,d,b)
}

