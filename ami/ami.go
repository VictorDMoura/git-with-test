package ami

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return (r.Altura * r.Largura)
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return c.Raio * c.Raio * math.Pi
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Altura + retangulo.Largura)
}
