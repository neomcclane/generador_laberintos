package laberinto

import (
	"fmt"
	"math/rand"
	"time"
)

type laberinto struct {
	Fila    int
	Columna int
	matriz  [][]string
}

func NewLaberinto(fila, columna int) *laberinto {
	var l *laberinto = new(laberinto)
	l.Fila = fila
	l.Columna = columna
	l.generarMatrizBase()
	return l
}

func (l *laberinto) Dibujar() {
	for fila := 0; fila < l.Fila; fila++ {
		for columna := 0; columna < l.Columna; columna++ {
			fmt.Print(l.matriz[fila][columna] + " ")
		}
		fmt.Println()
	}
}

func (l *laberinto) generarMatrizBase() {
	l.matriz = make([][]string, l.Fila)

	for fila := 0; fila < l.Fila; fila++ {
		l.matriz[fila] = make([]string, l.Columna)
		for columna := 0; columna < l.Columna; columna++ {
			l.matriz[fila][columna] = "#"
		}
	}
	l.generarEntrada()
}

func (l *laberinto) generarEntrada() {
	rand.Seed(time.Now().UTC().UnixNano())
	randFila := rand.Intn(l.Fila)
	randColumna := 0

	if randFila == 0 || randFila == (l.Fila-1) {
		rand.Seed(time.Now().UTC().UnixNano())
		randColumna = rand.Intn(l.Columna)
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
		randColumna = rand.Intn(2)
	}

	l.matriz[randFila][randColumna] = "e"
}
