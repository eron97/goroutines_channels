package main

import (
	"fmt"
	"time"
)

func myFunction(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go myFunction("goroutine")
	myFunction("main")

	/* output da main()
	goroutine
	main
	goroutine
	main
	main
	*/
}
