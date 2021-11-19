package MC2

import (
	"fmt"
	"math/rand"

	"concorrente/Sem1"
	"concorrente/Sem2"

	"golang.org/x/sync/semaphore"
)

func Producer1(buffer *[]float32, notEmpty, notFull *sem1.Semaphore) {
	for {
		d := rand.Float32()
		notFull.Wait()
		*buffer = append(*buffer, d)
		notEmpty.Signal()
		fmt.Printf("Produced %f\n", d)
	}
}

func Consumer1(buffer *[]float32, notEmpty, notFull *sem1.Semaphore) {
	for {
		notEmpty.Wait()
		d := (*buffer)[0]
		if len(*buffer) <= 1 {
			*buffer = []float32{}
		} else {
			*buffer = (*buffer)[1:]
		}
		notFull.Signal()
		fmt.Printf("Consumed %f\n", d)
	}
}

func Producer2(buffer *[]float32, notEmpty, notFull *sem2.Semaphore) {
	for {
		d := rand.Float32()
		notFull.Wait()
		*buffer = append(*buffer, d)
		notEmpty.Signal()
		fmt.Printf("Produced %f\n", d)
	}
}

func Consumer2(buffer *[]float32, notEmpty, notFull *sem2.Semaphore) {
	for {
		notEmpty.Wait()
		d := (*buffer)[0]
		if len(*buffer) <= 1 {
			*buffer = []float32{}
		} else {
			*buffer = (*buffer)[1:]
		}
		notFull.Signal()
		fmt.Printf("Consumed %f\n", d)
	}
}

func Producer3(buffer *[]float32, notEmpty, notFull *semaphore.Weighted) {
	for {
		d := rand.Float32()
		notFull.Acquire(nil, 1)
		*buffer = append(*buffer, d)
		notEmpty.Release(1)
		fmt.Printf("Produced %f\n", d)
	}
}

func Consumer3(buffer *[]float32, notEmpty, notFull *semaphore.Weighted) {
	for {
		notEmpty.Acquire(nil, 1)
		d := (*buffer)[0]
		if len(*buffer) <= 1 {
			*buffer = []float32{}
		} else {
			*buffer = (*buffer)[1:]
		}
		notFull.Release(1)
		fmt.Printf("Consumed %f\n", d)
	}
}
