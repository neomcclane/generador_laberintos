package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ejercicios/generador_laberintos/laberinto"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el numero de columnas: ")
	fila, _ := reader.ReadString('\n')
	fmt.Print("Ingrese el numero de filas: ")
	columna, _ := reader.ReadString('\n')

	iFila, _ := strconv.Atoi(strings.Split(fila, "\n")[0])
	iColumna, _ := strconv.Atoi(strings.Split(columna, "\n")[0])

	laberinto := laberinto.NewLaberinto(iFila, iColumna)
	laberinto.Dibujar()
}
