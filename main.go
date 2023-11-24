package main

import (
	"fmt"
	"sync"
)

/*

Criação de uma goroutine `contar` que envia números 1 a 10 para um channel.

Criação de uma goroutine anônima que recebe e imprime os números.

Uso de sync.WaitGroup é uma maneira de garantir que a func main() aguarde até que
todas as goroutines terminem seu trabalho antes de encerrar o programa.

*/

func contar(canal chan int, grupo *sync.WaitGroup) {
	defer grupo.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("enviando número para o channel")
		canal <- i
	}
	close(canal)
}

func main() {
	canal := make(chan int)
	var grupo sync.WaitGroup

	// inicia goroutine
	grupo.Add(1)
	go contar(canal, &grupo)

	// goroutine para imprimir os resultados
	go func() {
		for numero := range canal {
			fmt.Println("Contagem:", numero)
		}
	}()

	// espera todas goroutines terminarem
	grupo.Wait()

	/* output de main()
	enviando número para o channel
	enviando número para o channel
	Contagem: 1
	Contagem: 2
	enviando número para o channel
	enviando número para o channel
	Contagem: 3
	Contagem: 4
	enviando número para o channel
	enviando número para o channel
	Contagem: 5
	Contagem: 6
	enviando número para o channel
	enviando número para o channel
	Contagem: 7
	Contagem: 8
	enviando número para o channel
	enviando número para o channel
	Contagem: 9
	Contagem: 10
	*/
}
