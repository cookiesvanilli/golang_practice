//Создать структуры: Точка (x, y),
//Линия (точка начала, точка конца),
//Треугольник (три линии).
//Создать для каждой структуры переменную.
//Вычислить периметр каждой фигуры. Использовать интерфейс
package main

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
func (p Point) perimeter() float32 {
	return 0
}
func (l Line) perimeter() float32 {
	return 0 //TODO расчет расстояния между точками
}
func (t Triangle) perimeter() float32 {
	return t.l1.perimeter() + t.l2.perimeter() + t.l3.perimeter()
}
func main()  {

}






































