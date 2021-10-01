package canais2

func GeraMultiplos(c1 chan int, numMult, tam int) {
	for i := 1; i <= tam; i++ {
		c1 <- numMult * i
	}
}

func MesclaOrdenado(c1, c2, c3, cOut chan int, tam int) {
	var a, b, c int
	a = <-c1
	b = <-c2
	c = <-c3
	for i := 0; i < tam; i++ {

		if a < b && a < c {
			cOut <- a
			a = <-c1
			//prox posi do a
		} else if b < a && b < c {
			cOut <- b
			b = <-c2
			//prox posi do b
		} else if c < a && c < b {
			cOut <- c
			c = <-c3
			//prox posi do c
		}
	}
}

func Ex2() {
	println("== Ex. 2 ==")
	N := 10 // tamanho do canal
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	cOut := make(chan int)
	fim := make(chan struct{})

	go GeraMultiplos(c1, 2, N)
	go GeraMultiplos(c2, 3, N)
	go GeraMultiplos(c3, 5, N)

	go MesclaOrdenado(c1, c2, c3, cOut, N*3)

	go Consumir(cOut, fim, N*3)

	<-fim
}
