package laberinto

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const ESPACIO = "."
const PARED = "#"
const SALIDA = "S"
const ENTRADA = "E"

type laberinto struct {
	Fila    int
	Columna int
	matriz  [][]string

	pColumna       int
	pFila          int
	pInicioFila    int
	pInicioColumna int

	vPared  [][]int
	vCamino [][]int
}

func NewLaberinto(fila, columna int) *laberinto {
	var l *laberinto = new(laberinto)
	l.Fila = fila
	l.Columna = columna
	l.vPared = make([][]int, 0)
	l.vCamino = make([][]int, 0)
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
			l.matriz[fila][columna] = PARED
		}
	}
	l.generarEntrada()
}

func (l *laberinto) generarEntrada() {
	rand.Seed(time.Now().UTC().UnixNano())
	randFila := rand.Intn(l.Fila)
	randColumna := 0

	if randFila == 0 || randFila == (l.Fila-1) {
		for randColumna == 0 || randColumna == l.Columna-1 {
			rand.Seed(time.Now().UTC().UnixNano())
			randColumna = rand.Intn(l.Columna)
		}

	} else {
		rand.Seed(time.Now().UTC().UnixNano())
		randColumna = rand.Intn(2)

		if randColumna == 1 {
			randColumna = l.Columna - 1
		}
	}

	l.matriz[randFila][randColumna] = ENTRADA
	l.pFila = randFila
	l.pColumna = randColumna
	l.pInicioFila = l.pFila
	l.pInicioColumna = l.pColumna
	l.encontrarPuntoInicio()
}

func (l *laberinto) encontrarPuntoInicio() {
	if l.pColumna == 0 {
		l.pColumna = l.pColumna + 1

	} else if l.pColumna == l.Columna-1 {
		l.pColumna = l.pColumna - 1

	} else if l.pFila == 0 {
		l.pFila = l.pFila + 1

	} else if l.pFila == l.Fila-1 {
		l.pFila = l.pFila - 1
	}
	l.matriz[l.pFila][l.pColumna] = ESPACIO

	l.vPared = append(l.vPared, []int{l.pFila, l.pColumna})
	l.escavar()
}

func (l *laberinto) existeElementoVector(punto []int, vector [][]int) bool {
	resultado := false
	for _, fila := range vector {
		if fila[0] == punto[0] && fila[1] == punto[1] {
			resultado = true
		}
	}
	return resultado
}

func (l *laberinto) existeElementoCamino(punto []int) bool {
	resultado := false
	for _, fila := range l.vCamino {
		if fila[0] == punto[0] && fila[1] == punto[1] {
			resultado = true
		}
	}
	return resultado
}

func (l *laberinto) existeElementoPared(punto []int) bool {
	resultado := false
	for _, fila := range l.vPared {
		if fila[0] == punto[0] && fila[1] == punto[1] {
			resultado = true
		}
	}
	return resultado
}

func (l *laberinto) existeParedCamino(punto []int) bool {
	return !l.existeElementoPared(punto) && !l.existeElementoCamino(punto)
}

func (l *laberinto) escavar() {
	// fmt.Println(l.pFila, l.pColumna)
	vAuxiliarCamino := make([][]int, 0)

	if l.pColumna > 1 && l.existeParedCamino([]int{l.pFila, l.pColumna - 1}) && strings.EqualFold(l.matriz[l.pFila][l.pColumna-1], PARED) { // izquierda
		vAuxiliarCamino = append(vAuxiliarCamino, []int{l.pFila, l.pColumna - 1})

	}
	if l.pColumna < l.Columna-2 && l.existeParedCamino([]int{l.pFila, l.pColumna + 1}) && strings.EqualFold(l.matriz[l.pFila][l.pColumna+1], PARED) { // derecha
		vAuxiliarCamino = append(vAuxiliarCamino, []int{l.pFila, l.pColumna + 1})

	}
	if l.pFila > 1 && l.existeParedCamino([]int{l.pFila - 1, l.pColumna}) && strings.EqualFold(l.matriz[l.pFila-1][l.pColumna], PARED) { // arriba
		vAuxiliarCamino = append(vAuxiliarCamino, []int{l.pFila - 1, l.pColumna})

	}
	if l.pFila < l.Fila-2 && l.existeParedCamino([]int{l.pFila + 1, l.pColumna}) && strings.EqualFold(l.matriz[l.pFila+1][l.pColumna], PARED) { //abajo
		vAuxiliarCamino = append(vAuxiliarCamino, []int{l.pFila + 1, l.pColumna})

	}

	if len(vAuxiliarCamino) > 1 {
		time.Sleep(time.Nanosecond * 4889)
		rand.Seed(time.Now().UTC().UnixNano())
		tam := rand.Intn(len(vAuxiliarCamino)-1) + 1

		for i := 0; i < tam; i++ {
			pos := rand.Intn(len(vAuxiliarCamino))
			l.vCamino = append(l.vCamino, vAuxiliarCamino[pos])
			vAuxiliarCamino = append(vAuxiliarCamino[:pos], vAuxiliarCamino[pos+1:]...)
		}

		for j := 0; j < len(vAuxiliarCamino); j++ {
			l.vPared = append(l.vPared, vAuxiliarCamino[j])
		}

	} else if len(vAuxiliarCamino) == 1 {
		l.vCamino = append(l.vCamino, []int{vAuxiliarCamino[0][0], vAuxiliarCamino[0][1]})
		vAuxiliarCamino = append(vAuxiliarCamino[:0], vAuxiliarCamino[1:]...)
	}

	if len(l.vCamino) > 0 {
		rand.Seed(time.Now().UTC().UnixNano())
		pos := rand.Intn(len(l.vCamino))
		l.pFila = l.vCamino[pos][0]
		l.pColumna = l.vCamino[pos][1]
		l.vCamino = append(l.vCamino[:pos], l.vCamino[pos+1:]...)
		l.vPared = append(l.vPared, []int{l.pFila, l.pColumna})
		l.matriz[l.pFila][l.pColumna] = ESPACIO
		l.escavar()
	} else {
		l.generarSalida()
	}
}

func (l *laberinto) generarSalida() {

	switch {
	case l.pInicioFila == 0:

		for columna := 1; columna < l.Columna-2; columna++ {
			if strings.EqualFold(l.matriz[l.Fila-2][columna], ESPACIO) {
				l.matriz[l.Fila-1][columna] = SALIDA
				break
			}
		}

	case l.pInicioFila == l.Fila-1:

		for columna := 1; columna < l.Columna-2; columna++ {
			if strings.EqualFold(l.matriz[1][columna], ESPACIO) {
				l.matriz[0][columna] = SALIDA
				break
			}
		}

	case l.pInicioColumna == 0:

		for fila := 1; fila < l.Fila-2; fila++ {
			if strings.EqualFold(l.matriz[fila][l.Columna-2], ESPACIO) {
				l.matriz[fila][l.Columna-1] = SALIDA
				break
			}
		}
	case l.pInicioColumna == l.Columna-1:

		for fila := 1; fila < l.Fila-2; fila++ {
			if strings.EqualFold(l.matriz[fila][1], ESPACIO) {
				l.matriz[fila][0] = SALIDA
				break
			}
		}
	}
}
