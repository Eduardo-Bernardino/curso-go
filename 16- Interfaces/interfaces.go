package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.3f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

	fmt.Println(r)
	fmt.Println(c)
}
