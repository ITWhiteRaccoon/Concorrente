package main

func Gerar(c chan int, start, tam int) {
	for i := start; i < start+tam; i++ {
		c <- i
	}
}

func Mesclar(c1, c2, cOut1 chan int) {
	for {
		select {
		case dado := <-c1:
			cOut1 <- dado

		case dado := <-c2:
			cOut1 <- dado
		}
	}
}

func Consumir(c chan int, fim chan struct{}, tam int) {
	for i := 0; i < tam; i++ {
		dado := <-c
		println(dado)
	}
	fim <- struct{}{}
}

func Ex1a() {
	println("== Ex. 1a ==")
	N := 10
	c1 := make(chan int)
	c2 := make(chan int)
	cOut := make(chan int)
	fim := make(chan struct{})

	go Gerar(c1, 1000, N)
	go Gerar(c2, 2000, N)

	go Mesclar(c1, c2, cOut)

	go Consumir(cOut, fim, N*2)

	<-fim
}

func Ex1b() {
	println("== Ex. 1b ==")
	N := 10 // tamanho do canal
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	cOut1 := make(chan int)
	cOut2 := make(chan int)
	fim := make(chan struct{})

	go Gerar(c1, 1000, N)
	go Gerar(c2, 2000, N)
	go Gerar(c3, 3000, N)

	go Mesclar(c1, c2, cOut1)
	go Mesclar(cOut1, c3, cOut2)

	go Consumir(cOut2, fim, N*3)

	<-fim
}

func Ex1() {
	Ex1a()
	Ex1b()
	//1c. O aumento dos buffers possibilita que a geração dos números seja feita
	//concorrentemente à junção, sendo muito útil no caso de geradores mais
	//complexos que demorem para calcular um novo número
}
