//Eduardo C. Andrade e Julia A. Maia
package MC3

import (
	"fmt"
	"time"

	"concorrente/MCCSemaforo"
)

var readersB int

func Ex3B(NR, NW int) {
	readersB = 0
	mutex := MCCSemaforo.NewSemaphore(1)
	roomEmpty := MCCSemaforo.NewSemaphore(1)
	catraca := MCCSemaforo.NewSemaphore(1)

	for i := 0; i < NR; i++ {
		go ReaderB(mutex, roomEmpty, catraca)
	}
	for i := 0; i < NW; i++ {
		go WriterB(roomEmpty, catraca)
	}

	fin := make(chan struct{})
	<-fin
}

func WriterB(roomEmpty, catraca *MCCSemaforo.Semaphore) {
	for {
		catraca.Wait()

		roomEmpty.Wait()
		fmt.Println("critical section for writers")
		time.Sleep(1 * time.Second)
		catraca.Signal()

		roomEmpty.Signal()
	}
}

func ReaderB(mutex, roomEmpty, catraca *MCCSemaforo.Semaphore) {
	for {
		catraca.Wait()
		catraca.Signal()

		mutex.Wait()
		readersB += 1
		if readersB == 1 {
			roomEmpty.Wait()
		}
		mutex.Signal()

		fmt.Println("critical section for readers")
		time.Sleep(1 * time.Second)

		mutex.Wait()
		readersB -= 1
		if readersB == 0 {
			roomEmpty.Signal()
		}
		mutex.Signal()
	}
}
