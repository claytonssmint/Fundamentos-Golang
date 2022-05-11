package main

import (
	"fmt"
	"math"
)

type forma interface{
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f,\n", f.area())
}

type retangulo struct {
	altura float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// FORMA DE CALCULAR O CIRCULO É Pi, math tem funções matemáticas, para cálculos
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2) // Outra maneira de cálculos (c.raio * c.raio) 
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}