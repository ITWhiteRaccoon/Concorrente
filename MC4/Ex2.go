//Eduardo C. Andrade e Julia A. Maia
package MC4

import (
	"fmt"

	"concorrente/MCCSemaforo"
)

var (
	elves    int
	reindeer int
)

func Ex2() {
	elves = 0
	reindeer = 0
	santaSem := MCCSemaforo.NewSemaphore(0)
	reindeerSem := MCCSemaforo.NewSemaphore(0)
	elfTex := MCCSemaforo.NewSemaphore(1)
	mutex := MCCSemaforo.NewSemaphore(1)

	go Santa(santaSem, reindeerSem, mutex)
	for i := 0; i < 9; i++ {
		go Reindeer(santaSem, reindeerSem, mutex)
	}

	for i := 0; i < 9; i++ {
		go Elf(santaSem, elfTex, mutex)
	}

	fin := make(chan struct{})
	<-fin
}

func Santa(santaSem, reindeerSem, mutex *MCCSemaforo.Semaphore) {
	for {
		santaSem.Wait()
		mutex.Wait()
		if reindeer == 9 {
			fmt.Println("Santa is preparing the sleigh")
			for i := 0; i < 9; i++ {
				reindeerSem.Signal()
			}
		} else if elves == 3 {
			fmt.Println("Santa is helping the elves")
		}
		mutex.Signal()
	}
}

func Reindeer(santaSem, reindeerSem, mutex *MCCSemaforo.Semaphore) {
	mutex.Wait()
	reindeer += 1
	if reindeer == 9 {
		fmt.Println("The reindeers are waking up Santa")
		santaSem.Signal()
	}
	mutex.Signal()

	reindeerSem.Wait()
	fmt.Println("A reindeer is getting hitched")
}

func Elf(santaSem, elfTex, mutex *MCCSemaforo.Semaphore) {
	elfTex.Wait()
	mutex.Wait()
	elves += 1
	if elves == 3 {
		fmt.Println("The elves are waking up Santa")
		santaSem.Signal()
	} else {
		elfTex.Signal()
	}
	mutex.Signal()

	fmt.Println("An elf is getting help")

	mutex.Wait()
	elves -= 1
	if elves == 0 {
		elfTex.Signal()
	}
	mutex.Signal()
}
