package main

// Con los buffer puedes crear canales y asignales una capacidad.
// Si el canal no puede ser alamcenado en ninguna parte, quiere decir que necesitara un buffer
// Para cargar el buffer tiene que apuntar acia adentro del chanel
// Para librer el channel se necesita apuntar saliendo del chanel <- c
// Los buffers nos sirven para limitar la cantidad de procesos concurrentes que van a ser ejecutados al mismo tiempo [make(chan int, {bufferINT})]
//
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 12)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

// doSomething va a tardar mucho en ser ejecutado
func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("id %d started \n", i)
	time.Sleep(3 * time.Second)
	fmt.Printf("id %d finished \n", i)
	<-c

}
