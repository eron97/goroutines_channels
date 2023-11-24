package main

import (
	"fmt"
	"time"
)

func myFunction(canal chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		canal <- "goroutine"
	}
	close(canal)
}

func main() {
	canal := make(chan string)
	go myFunction(canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	/* output de main()
	goroutine
	goroutine
	goroutine
	*/
}
