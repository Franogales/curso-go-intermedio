package main

import "fmt"

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}

// Aqui usamos el la flechita apuntado a la izquierda, posicionada al lado derecho de la parabra chan para que sea exclusivamente de escritura
func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c) // cuando esto se llama, se bloquea el canal (ya no comunica)
}

// Aqui usamos la fleshita a la izquierda de la palabra chan para exclusiva lectura
func Double(in <-chan int, out chan int) {
	for val := range in {
		out <- 2 * val
	}
	close(out)
}

func Print(c chan int) {
	for val := range c {
		fmt.Println(val)
	}
}
