//Создать структуры: Точка (x, y),
//Линия (точка начала, точка конца),
//Треугольник (три линии).
//Создать для каждой структуры переменную.
//Вычислить периметр каждой фигуры. Использовать интерфейс
package main

import (
	"fmt"

)

type Figures interface {
	perimeter()
}
type Point1 struct {
 	x, y int
}
type Point2 struct {
	x, y int
}
type Line1 struct {
	Point1
	Point2
}
type Line2 struct {
	Point1
	Point2
}
type Line3 struct {
	Point1
	Point2
}
type Triangle struct {
	Line1
	Line2
	Line3
}

//Triangle
func (t *Triangle) perimeter(){
	a:= Line1{Point1{x: 3, y: 3}, Point2{x: 5, y: 4}}
	b:= Line2{Point1{x: 4, y: 3}, Point2{x: 3, y: 3}}
	c:= Line3{Point1{x: 5, y: 3}, Point2{x: 3, y: 7}}

}
func main()  {

}






































