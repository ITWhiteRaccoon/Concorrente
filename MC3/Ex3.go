//Eduardo C. Andrade e Julia A. Maia
package MC3

import (
	"fmt"
	"time"

	"concorrente/MCCSemaforo"
)

var readers int

func Ex3(NR, NW int) {
	readers = 0
	mutex := MCCSemaforo.NewSemaphore(1)
	roomEmpty := MCCSemaforo.NewSemaphore(1)

	for i := 0; i < NR; i++ {
		go Reader(mutex, roomEmpty)
	}
	for i := 0; i < NW; i++ {
		go Writer(roomEmpty)
	}

	fin := make(chan struct{})
	<-fin
}

func Writer(roomEmpty *MCCSemaforo.Semaphore) {
	for {
		roomEmpty.Wait()

		fmt.Println("critical section for writers")
		time.Sleep(1 * time.Second)

		roomEmpty.Signal()
	}
}

func Reader(mutex, roomEmpty *MCCSemaforo.Semaphore) {
	for {
		mutex.Wait()
		readers += 1
		if readers == 1 {
			roomEmpty.Wait()
		}
		mutex.Signal()

		fmt.Println("critical section for readers")
		time.Sleep(1 * time.Second)

		mutex.Wait()
		readers -= 1
		if readers == 0 {
			roomEmpty.Signal()
		}
		mutex.Signal()
	}
}
