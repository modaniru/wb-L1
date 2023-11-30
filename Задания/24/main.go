package main

import (
	"fmt"
	"math"
)

//К переменным можем обращаться только в этом пакете
type Point struct{
	x float64
	y float64
}

func NewPoint(x, y float64) *Point{
	return &Point{x: x, y: y}
}

func (p *Point) GetX() float64{
	return p.x
}

func (p *Point) GetY() float64{
	return p.y
}

//Метод для вычисления растояние от точки до точки
func (p *Point) Distance(p1 *Point) float64{
	return math.Sqrt(math.Pow(math.Abs(p.x - p1.x), 2) + math.Pow(math.Abs(p.y - p1.y), 2))
}

func main(){
	p1 := NewPoint(4, 3)
	p2 := NewPoint(0, 0)
	fmt.Println(p1.Distance(p2))
	fmt.Println(p2.Distance(p1))

}