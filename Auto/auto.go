//Существует набор структур: Колеса, Двигатель, Корпус. Колеса могут быть разные,
//Двигатели тоже разные (например, 4 цилиндра, 8 цилиндров, дизель, бензиновый).
//Корпуса могут быть седан, хэтчбэк, универсал, можно добавить свои.
//Собрать из этого набора несколько разных видов машин.
package main

import "fmt"

type Automobile interface {
	season() string
	cylinder() string
	model() string
}
type Wheels struct {
    structure string
}
type Engine struct {
	structure string
}
type Corpus struct {
	structure string
}
type MyCar struct {
	audi []Automobile
}
func (c *Wheels) season() string {
	return "Wheels: winter"
}
func (e *Engine) cylinder() string {
	return "Cylinder: 8 cylinders"
}
func (m *Corpus) model() string {
	return "Model: sedan"
}
func (car *MyCar) season() string {
	var season string
	for _, s := range car.audi {
		season += s.season()
	}
	return season
}
func main() {
	c := Wheels{structure: "Wheels: winter"}
	fmt.Println(c.season())
	e := Engine{structure: "Cylinder: 8 cylinders"}
	fmt.Println(e.cylinder())
	m := Corpus{structure: "Model: sedan"}
	fmt.Println(m.model())
}