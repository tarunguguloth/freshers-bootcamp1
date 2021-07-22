package main

import (
	"fmt"
	"sync"
)

var(
	mutex sync.Mutex
	balance int
)
func deposit(value int, wg *sync.WaitGroup){
	mutex.Lock()
	balance += value
	fmt.Printf("Deposited %d amount and new current balance %d\n",value, balance)
	mutex.Unlock()
	wg.Done()
}
func withdraw(value int, wg *sync.WaitGroup){
	mutex.Lock()
	if(balance<value){
		fmt.Println("There in no enough money")
	} else{

	balance -= value
	fmt.Printf("Withdrawing %d amount and new current balance %d\n",value, balance)
	}
	mutex.Unlock()
	wg.Done()
}
func main() {
	balance = 500
	var wg sync.WaitGroup
	wg.Add(2)

	go withdraw(500, &wg)
	go deposit(1200, &wg)
	wg.Wait()

	fmt.Printf("Final Balance %d\n", balance)
	}