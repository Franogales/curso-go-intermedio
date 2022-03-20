package main

import (
	"fmt"
	"time"
)

// Esto nos ayuda para saber que chanel se ejecuta y no es de una manera secuencial
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second
	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		// el select nos ayudara a ver exactamente cual fue el canal activado.
		select {
		case channelMessage1 := <-c1:
			fmt.Println(channelMessage1)
		case channelMessage2 := <-c2:
			fmt.Println(channelMessage2)

		}

	}

	// de otra manera tendriamos que esperar a que c1 (que tarda 4 segundos) a que termine para imprimir
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
