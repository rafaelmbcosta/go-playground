package main

import "math"

// Letra MAÍUSCULA é Publico (visível dentro e fora do PACOTE).
// Letra minúscula é Privado (visível apenas dentro do PACOTE).

// Example --> public
// example or _Example --> private

// Line represents a coordinate
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// responsavel por calcular a distancia
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
