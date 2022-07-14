package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

func printArea(s shape){
	fmt.Println("the shape size is : " + fmt.Sprintf("%f", s.getArea()))
}

type square struct{
	length int
}

func (s square) getArea() float64{
	return float64(s.length * s.length)
}

type round struct {
	radius int
}

func (r round) getArea() float64 {
	return float64(r.radius) * float64(3.1415)
}

func main (){
	s := square {10}
	r := round {5}

	printArea(s)
	printArea(r)
}

