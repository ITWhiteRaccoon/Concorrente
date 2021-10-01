//Eduardo C. Andrade e Julia A. Maia
package canais3

import (
	"fmt"
	"strconv"
)

func philosopherB(id int, forkLeft, forkRight, table chan struct{}) {
	for {
		fmt.Println(strconv.Itoa(id) + " senta")
		select {
		case t := <-table: //Caso não bloqueado (ninguém comendo), bloqueia a mesa e come
			l := <-forkLeft
			r := <-forkRight
			fmt.Println(strconv.Itoa(id) + " pegou esquerda")
			fmt.Println(strconv.Itoa(id) + " pegou direita")
			fmt.Println(strconv.Itoa(id) + " come")
			forkLeft <- l
			forkRight <- r
			table <- t
		default: //Caso a mesa esteja sendo usada, levanta
			fmt.Println(strconv.Itoa(id) + " levanta e pensa")
			break
		}
	}
}

func runB() {
	var forkChannels [FORKS]chan struct{}
	//Como sugerido pelo prof., criamos um canal que possibilita o bloqueio da mesa enquanto um filósofo tem dois garfos
	tableChannel := make(chan struct{}, 1)
	tableChannel <- struct{}{}
	for i := 0; i < FORKS; i++ {
		forkChannels[i] = make(chan struct{}, 1)
		forkChannels[i] <- struct{}{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i))
		go philosopherB(i, forkChannels[i], forkChannels[(i+1)%PHILOSOPHERS], tableChannel)
	}
	var blq = make(chan struct{})
	<-blq
}
