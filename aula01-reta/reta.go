package main

import "math"

//inicializando qualquer coisa com letra Maisucula ele eh publico automaticamente
//inicializando qualquer coisa com letra minuscula ele eh privado

//Por exemplo de nomes...
// Ponto - gerara algo publico
// ponto ou _Ponto - gerara algo privado

//Coisas publicas devem ter um comentario em cima obrigatorio descrevendo o escopo

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia eh responsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
