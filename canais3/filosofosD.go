//Eduardo C. Andrade e Julia A. Maia
package canais3

import (
	"fmt"
	"strconv"
)

func philosopherD(id int, forkLeft, forkRight chan struct{}) {
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

func runD() {
	//Usando a quantidade de filósofos, criamos dois garfos para cada
	const availableForks = PHILOSOPHERS * 2
	var forkChannels [availableForks]chan struct{}
	for i := 0; i < availableForks; i++ {
		forkChannels[i] = make(chan struct{}, 1)
		forkChannels[i] <- struct{}{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i))
		//Alterando o método original, cada filósofo recebe um garfo exclusivo, eliminando a chance de deadlock
		go philosopherD(i, forkChannels[i*2], forkChannels[(i*2)+1])
	}
	var blq = make(chan struct{})
	<-blq
}
