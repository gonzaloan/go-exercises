package main

import (
	"fmt"
	"sync"
)

// Just 1 1 DEposit -> writing (Race Condition
// But N people can be depositing

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	balance += amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	/*
		Se bloquean las escrituras, pero no las lecturas. Eso significa que pueden haber
		diferentes lectores usando Balance() sin que sean bloqueados, pero lo sescritores siguen funcionando igual.
	*/
	return b
}

/*
En esencia, RLock de RWLock garantiza una secuencia de lecturas en donde el valor que lees no se verá alterado por
nuevos escritores, a diferencia de no usar nada.
*/
func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}
