//Существует набор структур: Колеса, Двигатель, Корпус. Колеса могут быть разные,
//Двигатели тоже разные (например, 4 цилиндра, 8 цилиндров, дизель, бензиновый).
//Корпуса могут быть седан, хэтчбек, универсал, можно добавить свои.
//Собрать из этого набора несколько разных видов машин.
package main

import (
	"fmt"
)

type Tracks struct {}
type Wheels struct {
    numberWheels int
}
type Engine struct {
	cylinderNumb int
	fuel string
}
type Corpus struct {
	typeOfCar string
}
type Car struct {
	w Wheels
	e Engine
	c Corpus
}
type Tractor struct {
	t Tracks
	e Engine
	c Corpus
}
func (t Tracks) madeof() {
	fmt.Println("Tracks")
}
func (w Wheels) madeof() {
	fmt.Println("Wheels: ", w.numberWheels)
}
func (e Engine) madeof()  {
	fmt.Println("Cylinders: ", e.cylinderNumb)
	fmt.Println("Fuel: ", e.fuel)
}
func (m Corpus) madeof()  {
	fmt.Println("Model: ", m.typeOfCar)
}
func (c Car) madeof() {
	c.w.madeof()
	c.e.madeof()
	c.c.madeof()
}
func (t Tractor) madeof() {
	t.t.madeof()
	t.e.madeof()
	t.c.madeof()
}
func main() {
	w := Wheels{numberWheels: 4}
	e := Engine{cylinderNumb: 8, fuel: "diesel"}
	c := Corpus{typeOfCar: "sedan"}
	car := Car{w, e, c}
	car.madeof()
	t := Tracks{}
	tractor := Tractor{t, e, c}
	tractor.madeof()
}