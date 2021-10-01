package main

func Ping(c1, c2, fim chan struct{}, tam, channel int) {
	for i := 0; i < tam; i++ {
		x := <-c1
		println(channel, i)
		c2 <- x
	}
	fim <- struct{}{}
}

func Pong(c1, c2, fim chan struct{}, tam, channel int) {
	for i := 0; i < tam; i++ {
		x := <-c1
		println(channel, i)
		if i < tam-1 {
			c2 <- x
		}
	}
	fim <- struct{}{}
}

func Ex4() {
	Ex4a()
	Ex4b()
}

func Ex4a() {
	println("== Ex. 4a ==")
	N := 10
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	fim := make(chan struct{}, 1) // segundo parametro: tamanho do buffer do fim

	go Ping(c1, c2, fim, N, 1)
	go Pong(c2, c1, fim, N, 2)
	c1 <- struct{}{}

	<-fim
	<-fim
}

func Ex4b() {
	println("== Ex. 4b ==")
	N := 3
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	fim := make(chan struct{}, 2) // segundo parametro: tamanho do buffer do fim

	go Ping(c1, c2, fim, N, 1)
	go Ping(c2, c3, fim, N, 2)
	go Pong(c3, c1, fim, N, 3)
	c1 <- struct{}{}

	<-fim
	<-fim
	<-fim
}
