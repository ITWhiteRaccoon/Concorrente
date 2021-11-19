//Eduardo C. Andrade e Julia A. Maia
package MC3

import (
	"fmt"
	"math/rand"

	"concorrente/Barr"
)

var matriz [][]int
var N int

func Ex1(n int) {
	N = n
	matriz = make([][]int, N)
	for i := 0; i < N; i++ {
		matriz[i] = make([]int, N)
		for j := 0; j < N; j++ {
			matriz[i][j] = rand.Intn(N * 2)
		}
	}

	barreira := Barr.NewBarrier(N*N + 1)
	go ImprimirMatriz(barreira)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			go Processo(barreira, i, j)
		}
	}

	fin := make(chan struct{})
	<-fin
}

func Processo(barreira *Barr.Barrier, x, y int) {
	for i := 0; i < N; i++ {
		media := CalcularMedia(x, y)
		barreira.Arrive()
		EscreverMedia(media, x, y)
		barreira.Leave()
	}
}

func CalcularMedia(x, y int) int {
	var vizinhos []int
	if x > 0 {
		vizinhos = append(vizinhos, matriz[x-1][y])
	}
	if x < N-1 {
		vizinhos = append(vizinhos, matriz[x+1][y])
	}
	if y > 0 {
		vizinhos = append(vizinhos, matriz[x][y-1])
	}
	if y < N-1 {
		vizinhos = append(vizinhos, matriz[x][y+1])
	}
	soma := 0
	for i := 0; i < len(vizinhos); i++ {
		soma += vizinhos[i]
	}
	return soma / len(vizinhos)
}

func EscreverMedia(valor, x, y int) {
	matriz[x][y] = valor
}

func ImprimirMatriz(barreira *Barr.Barrier) {
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				fmt.Printf(" %d ", matriz[i][j])
			}
			fmt.Println()
		}
		barreira.Arrive()
		barreira.Leave()
	}
}
