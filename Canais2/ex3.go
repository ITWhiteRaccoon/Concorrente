package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999999999)
	}
	return slice
}

//Recebe uma lista de inteiros, conta os primos, e devolve quantos são
func contaPrimosConc(fim chan int, inteiros []int) {
	result := 0
	c := make(chan bool)
	for _, x := range inteiros {
		go isPrimeConc(c, x)
	}
	for i := 0; i < len(inteiros); i++ {
		prime := <-c
		if prime {
			result++
		}
	}
	fim <- result
}

func isPrimeConc(c chan bool, p int) {
	if p%2 == 0 {
		c <- false
	}
	for i := 3; i*i <= p; i += 2 {
		if p%i == 0 {
			c <- false
		}
	}
	c <- true
}

const N = 2000

func Ex3() {
	println("== Ex. 3 ==")
	slice := generateSlice(N)
	fim := make(chan int)
	go contaPrimosConc(fim, slice)
	p := <-fim
	fmt.Println("n primos : ", p)
}
