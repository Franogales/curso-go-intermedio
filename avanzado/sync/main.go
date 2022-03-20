package main

// Para revisar si tiene peligro de carrera se usa el go build --race main.go
import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock() //se bloquea el programa y tiene que esperar hasta que el que tiene lock lo desbloquee. La siguiente goroutine tiene que esperar a que
	// este lock llegue a unlock -(condicion de carrera)
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup //para bloquear el programa
	var lock sync.Mutex   //para evitar que el deposito este usando la misma variable en diferentes goroutnes
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())

}
