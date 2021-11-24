//Eduardo C. Andrade e Julia A. Maia
package MC4

import (
	"fmt"

	"concorrente/MCCSemaforo"
)

var (
	customers     int
	customerLimit int
	walkedAway    int
	served        int
	barberClose   chan struct{}
)

func Ex1(custLimit, nCust int) {
	barberClose = make(chan struct{}, nCust)
	customers = 0
	walkedAway = 0
	served = 0
	customerLimit = custLimit
	mutex := MCCSemaforo.NewSemaphore(1)
	customer := MCCSemaforo.NewSemaphore(0)
	barber := MCCSemaforo.NewSemaphore(0)

	go Barber(customer, barber)
	for i := 0; i < nCust; i++ {
		go Customer(mutex, customer, barber)
	}

	for i := 0; i < nCust; i++ {
		<-barberClose
	}

	fmt.Printf("\n===END OF DAY===\n")
	fmt.Printf("N of customers served: %d\n", served)
	fmt.Printf("N of customers that walked away: %d\n", walkedAway)

}

func Customer(mutex, customer, barber *MCCSemaforo.Semaphore) {
	mutex.Wait()
	if customers == customerLimit+1 {
		walkedAway += 1
		mutex.Signal()
		barberClose <- struct{}{}
		fmt.Println("Customer walked away")
		return
	}
	customers += 1
	mutex.Signal()
	fmt.Println("Customer is waiting")

	customer.Signal()
	barber.Wait()
	fmt.Println("Customer sat on barber's chair")

	mutex.Wait()
	customers -= 1
	served += 1
	mutex.Signal()
	barberClose <- struct{}{}
}

func Barber(customer, barber *MCCSemaforo.Semaphore) {
	for {
		customer.Wait()
		fmt.Println("Barber got woken up")
		barber.Signal()
		fmt.Println("Barber cut hair")
		fmt.Println("Barber is sleeping")
	}
}
