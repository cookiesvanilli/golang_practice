//Существует набор структур: Колеса, Двигатель, Корпус. Колеса могут быть разные,
//Двигатели тоже разные (например, 4 цилиндра, 8 цилиндров, дизель, бензиновый).
//Корпуса могут быть седан, хэтчбек, универсал, можно добавить свои.
//Собрать из этого набора несколько разных видов машин.
package main

import (
	"fmt"
)

type AirplaneWings struct {}
type BaggageHold struct {}
type Tracks struct {}
type Wheels struct {
    numberWheels int
}
type Engine struct {
	cylinderNumb int
	fuel string
}
type Corpus struct {
	typeOfTransport string
}
type Airplane struct {
	a AirplaneWings
	b BaggageHold
	w Wheels
	e Engine
	c Corpus
}
type Car struct {
	w Wheels
	e Engine
	c Corpus
	b BaggageHold

}
type Tractor struct {
	t Tracks
	e Engine
	c Corpus
}
func (a AirplaneWings) madeof() {
	fmt.Println("Airplane wings")
}
func (b BaggageHold) madeof() {
	fmt.Println("Baggage")
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
	fmt.Println("Type of Transport: ", m.typeOfTransport)
}
func (a Airplane) madeof() {
	a.a.madeof()
	a.b.madeof()
	a.e.madeof()
	a.w.madeof()
	a.c.madeof()
}
func (c Car) madeof() {
	c.w.madeof()
	c.e.madeof()
	c.c.madeof()
	c.b.madeof()
}
func (t Tractor) madeof() {
	t.t.madeof()
	t.e.madeof()
	t.c.madeof()
}
func main() {
	fmt.Println("Car =======")
	w := Wheels{numberWheels: 4}
	e := Engine{cylinderNumb: 8, fuel: "diesel"}
	c := Corpus{typeOfTransport: "sedan"}
	b := BaggageHold{}
	car := Car{w, e, c, b}
	car.madeof()
	fmt.Println("Tractor =======")
	t := Tracks{}
	tc := Corpus{typeOfTransport: "bulldozer"}
	tractor := Tractor{t, e, tc}
	tractor.madeof()
	fmt.Println("Airplane =======")
	a := AirplaneWings{}
	h := BaggageHold{}
	aw := Wheels{numberWheels: 2}
	ae := Engine{cylinderNumb: 6, fuel: "petrol"}
	ac := Corpus{typeOfTransport: "boat-shaped"}
	plane := Airplane{a, h, aw, ae, ac}
	plane.madeof()

}