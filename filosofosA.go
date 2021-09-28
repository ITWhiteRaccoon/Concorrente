package main

import (
	"fmt"
	"strconv"
)

func philosopherA(id int, forkLeft, forkRight chan struct{}) {
	for {
		fmt.Println(strconv.Itoa(id) + " senta")
		<-forkLeft
		fmt.Println(strconv.Itoa(id) + " pegou esquerda")
		<-forkRight
		fmt.Println(strconv.Itoa(id) + " pegou direita")
		fmt.Println(strconv.Itoa(id) + " come")
		forkLeft <- struct{}{}
		forkRight <- struct{}{}
		fmt.Println(strconv.Itoa(id) + " levanta e pensa")
	}
}

func runA() {
	var forkChannels [FORKS]chan struct{}
	for i := 0; i < FORKS; i++ {
		forkChannels[i] = make(chan struct{}, 1)
		forkChannels[i] <- struct{}{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS - 1); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i))
		go philosopherA(i, forkChannels[i], forkChannels[(i+1)%PHILOSOPHERS])
	}
	//Analisando o problema, concluímos que trocar a ordem dos canais simularia a
	//troca da ordem em que este filósofo pega os garfos.
	go philosopherA(PHILOSOPHERS-1, forkChannels[0], forkChannels[PHILOSOPHERS-1])
	var blq = make(chan struct{})
	<-blq
}
