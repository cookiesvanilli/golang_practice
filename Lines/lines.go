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
type PointXY struct {
 	x1, y1 float64
}
type Line struct {
	PointXY
	x2, y2 float64
}
type Triangle struct {
	Line
	x3, y3 float64
}
func points(x1, y1, x2, y2, x3, y3  float64) float64 {
	a := math.Sqrt( math.Pow((x2 - x1), 2) + math.Pow((y2-y1), 2))
	b := math.Sqrt( math.Pow((x3 - x2), 2) + math.Pow((y3-y2), 2))
	c := math.Sqrt( math.Pow((x1 - x3), 2) + math.Pow((y1-y3), 2))
	return (a + b + c)/2
}
func shapes(list ...Figures) {
	for _, a := range list {
		a.perimeter()
	}
}
//Triangle
func (t *Triangle) perimeter(){
	x := points(t.x1, t.y1, t.x2, t.y2, t.x3, t.y3)
	fmt.Println(x)
}
func main()  {
	/*l:= new(Triangle)
	l.x1 = 1
	l.y1 = 2
	l.x2 = 3
	l.y2 = 2
	l.x3 = 5
	l.y3 = 3*/
	s:= &Triangle{
		Line: Line{PointXY{
			x1: 3,
			y1: 5,
		},3,2},
		x3:   6,
		y3:   2,
	}
	shapes(s)
}






































