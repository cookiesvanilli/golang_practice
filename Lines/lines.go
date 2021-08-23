//Создать структуры: Точка (x, y),
//Линия (точка начала, точка конца),
//Треугольник (три линии).
//Создать для каждой структуры переменную.
//Вычислить периметр каждой фигуры. Использовать интерфейс
package main

import (
	"fmt"
	"math"
)

type Figures interface {
	perimeter()
}
type Point struct {
 	x, y int
}
type Line struct {
	p1 Point
	p2 Point
}
type Triangle struct {
	l1 Line
	l2 Line
	l3 Line
}
func (p Point) perimeter() int {
	return 0
}
func (l Line) perimeter() int {
	resX := l.p2.x - l.p1.y
	rX := math.Pow(float64(resX), 2)
	resY := l.p2.x - l.p1.y
	rY := math.Pow(float64(resY), 2)
	return int(math.Sqrt(rX + rY)) //TODO расчет расстояния между точками
}
func (t Triangle) perimeter() int {
	return t.l1.perimeter() + t.l2.perimeter() + t.l3.perimeter()
}
func main()  {
	t := Triangle{l1: Line{p1: Point{9,5}, p2: Point{5,2}},
				  l2: Line{p1: Point{7,7}, p2: Point{3,6}},
				  l3: Line{p1: Point{4,9}, p2: Point{2,7}}}
	t.perimeter()
	fmt.Println(t.perimeter())
}






































