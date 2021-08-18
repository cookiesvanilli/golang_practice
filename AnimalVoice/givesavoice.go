package main

import (
	"fmt"
	"reflect"
)

type Givesavoice interface {
	sayHello() string
}
type Cat struct {
}
type Dog struct {
}
type Bee struct {
}
func (c Cat) sayHello() string { //Cat
	return "MEOW"
}
func (d Dog) sayHello()  string { //Dog
	return "WOFF"
}
func (b Bee) sayHello() string  { //Bird
	return "BZZ"
}
func animalVoice(list ...Givesavoice) {
	for _, a := range list {
		fmt.Println(reflect.TypeOf(a).Name(), a.sayHello())
	}
}
func main() {
	c := Cat{}
	d := Dog{}
	b := Bee{}
	animalVoice(c,d,b)
}
