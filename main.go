/*
	Juego de la vida de Conway implentado en lenguaje Go
	2025 Felipe González Hernández
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXGENERATIONS int = 100

var size int

/*
Inicializa y genera aleatoriamente la prímera generación del juego
*/
func createUniverse() [][]int {
	universe := make([][]int, size)
	for i := 0; i < size; i++ {
		universe[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			universe[i][j] = rand.Intn(2)
		}
	}
	return universe
}

/*
Muestra por consola una generación del juego de la vida
Celda muerta = espacio ' '
Celda viva = letra o mayúscula 'O'
*/
func printGeneration(u [][]int) {
	states := " O"
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%c", states[u[i][j]])
		}
		fmt.Println()
	}
}

/*
Devuelve una matriz con la siguiente generación siguiendo las reglas del juego de la vida de Conway
*/
func nextGen(oldU [][]int) [][]int {
	var iMinus1, iPlus1 int
	var jMinus1, jPlus1 int

	nextU := make([][]int, size)
	for i := 0; i < size; i++ {
		nextU[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		if i-1 < 0 {
			iMinus1 = size - 1
		} else {
			iMinus1 = i - 1
		}
		if i+1 == size {
			iPlus1 = 0
		} else {
			iPlus1 = i + 1
		}
		for j := 0; j < size; j++ {
			if j-1 < 0 {
				jMinus1 = size - 1
			} else {
				jMinus1 = j - 1
			}
			if j+1 == size {
				jPlus1 = 0
			} else {
				jPlus1 = j + 1
			}
			neighbors := 0
			if oldU[iMinus1][jMinus1] == 1 {
				neighbors++
			}
			if oldU[iMinus1][j] == 1 {
				neighbors++
			}
			if oldU[iMinus1][jPlus1] == 1 {
				neighbors++
			}
			if oldU[i][jMinus1] == 1 {
				neighbors++
			}
			if oldU[i][jPlus1] == 1 {
				neighbors++
			}
			if oldU[iPlus1][jMinus1] == 1 {
				neighbors++
			}
			if oldU[iPlus1][j] == 1 {
				neighbors++
			}
			if oldU[iPlus1][jPlus1] == 1 {
				neighbors++
			}

			if oldU[i][j] == 1 { //Comprueba si una celda muere por aburrimiento o sobrepoblación
				if neighbors < 2 || neighbors > 3 {
					nextU[i][j] = 0
				} else {
					nextU[i][j] = 1
				}
			} else if neighbors == 3 { //Comprueba si una celda muerta deve revivir
				nextU[i][j] = 1
			} else {
				nextU[i][j] = 0
			}
		}
	}
	return nextU
}

// Devuelve el número de celdas del universo que están vivas en la generación actual
func getAlive(u [][]int) int {
	cont := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if u[i][j] == 1 {
				cont++
			}
		}
	}
	return cont
}
func main() {
	_, err := fmt.Scanln(&size)
	if err != nil {
		return
	}
	universe := createUniverse()
	for i := 1; i < MAXGENERATIONS+1; i++ {
		fmt.Printf("Generation #%d\n", i)
		fmt.Printf("Alive: %d\n", getAlive(universe))
		printGeneration(universe)
		universe = nextGen(universe)
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
	}
}
