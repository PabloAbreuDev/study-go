package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()   // Bloqueia o acesso
	c.count++     // Incrementa o contador
	c.mu.Unlock() // Libera o acesso
}

func sendData(ch1 chan<- string, ch2 chan<- string) {
	time.Sleep(2 * time.Second)
	ch1 <- "Data from channel 1"

	time.Sleep(1 * time.Second)
	ch2 <- "Data from channel 2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendData(ch1, ch2)

	// Usando o select para esperar por dados em múltiplos canais
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait() // Aguarda todas as goroutines terminarem
	fmt.Println("Final counter:", counter.count)

}
