package main

import (
	"fmt"
	"strconv"
)

func philosopherC(id int, forkLeft, forkRight chan struct{}) {
	for {
		fmt.Println(strconv.Itoa(id) + " senta")
		select {
		case l := <-forkLeft:
			fmt.Println(strconv.Itoa(id) + " pegou esquerda")
			select {
			case r := <-forkRight:
				forkLeft <- l
				forkRight <- r
				fmt.Println(strconv.Itoa(id) + " pegou direita")
				fmt.Println(strconv.Itoa(id) + " come")
			default:
				forkLeft <- l
				fmt.Println(strconv.Itoa(id) + " levanta e pensa")
				break
			}
		default:
			fmt.Println(strconv.Itoa(id) + " levanta e pensa")
			break
		}
	}
}

func runC() {
	var forkChannels [FORKS]chan struct{}
	for i := 0; i < FORKS; i++ {
		forkChannels[i] = make(chan struct{}, 1)
		forkChannels[i] <- struct{}{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i))
		go philosopherC(i, forkChannels[i], forkChannels[(i+1)%PHILOSOPHERS])
	}
	var blq = make(chan struct{})
	<-blq
}
