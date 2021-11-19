//Eduardo C. Andrade e Julia A. Maia
package MC2

import (
	"fmt"

	sem1 "concorrente/Sem1"
)

const PHILOSOPHERS int = 5

func RunA() {
	var forks [PHILOSOPHERS]*sem1.Semaphore
	for i := 0; i < PHILOSOPHERS; i++ {
		forks[i] = sem1.NewSemaphore(1)
	}
	for i := 0; i < PHILOSOPHERS-1; i++ {
		go philosopherA(i, forks[i], forks[(i+1)%PHILOSOPHERS])
	}
	go philosopherA(PHILOSOPHERS-1, forks[0], forks[PHILOSOPHERS-1])
	fin := make(chan struct{})
	<-fin
}

func philosopherA(id int, forkLeft, forkRight *sem1.Semaphore) {
	for {
		fmt.Printf("%d is thinking\n", id)
		forkLeft.Wait()
		forkRight.Wait()
		fmt.Printf("%d eat\n", id)
		forkLeft.Signal()
		forkRight.Signal()
	}
}

func RunB() {
	var forks [PHILOSOPHERS]*sem1.Semaphore
	table := sem1.NewSemaphore(1)
	for i := 0; i < PHILOSOPHERS; i++ {
		forks[i] = sem1.NewSemaphore(1)
	}
	for i := 0; i < PHILOSOPHERS; i++ {
		go philosopherB(i, forks[i], forks[(i+1)%PHILOSOPHERS], table)
	}
	fin := make(chan struct{})
	<-fin
}

func philosopherB(id int, forkLeft, forkRight, table *sem1.Semaphore) {
	for {
		fmt.Printf("%d is thinking\n", id)
		table.Wait()
		forkLeft.Wait()
		forkRight.Wait()
		fmt.Printf("%d eat\n", id)
		forkLeft.Signal()
		forkRight.Signal()
		table.Signal()
	}
}
