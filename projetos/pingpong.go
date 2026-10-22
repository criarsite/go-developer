package projetos

import (
	"fmt"
	"sync"
)

func Alternar() {
	canal := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			canal <- "ping"
			fmt.Println(<-canal)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			fmt.Println(<-canal)
			canal <- "pong"
		}
	}()

	wg.Wait()
}
