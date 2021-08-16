//Существует набор структур: Колеса, Двигатель, Корпус. Колеса могут быть разные,
//Двигатели тоже разные (например, 4 цилиндра, 8 цилиндров, дизель, бензиновый).
//Корпуса могут быть седан, хэтчбек, универсал, можно добавить свои.
//Собрать из этого набора несколько разных видов машин.
package main

import "fmt"

type Automobile interface {
	madeof()
}
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
func (w *Wheels) madeof() {
	fmt.Println("Wheels: ", w.numberWheels)
}
func (e *Engine) madeof()  {
	fmt.Println("Cylinders: ", e.cylinderNumb)
	fmt.Println("Fuel: ", e.fuel)
}
func (m *Corpus) madeof()  {
	fmt.Println("Model: ", m.typeOfCar)
}
func myCar(list ...Automobile) {
	for _, a := range list {
		a.madeof()
	}
}
func main() {
	w := &Wheels{numberWheels: 4}
	e := &Engine{cylinderNumb: 8, fuel: "diesel"}
	m := &Corpus{typeOfCar: "sedan"}
	car := []Automobile{w, e, m}
	myCar(car...)
}