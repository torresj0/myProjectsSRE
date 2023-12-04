package main

//duas goroutines se comunicam e sincronizam sua execucao
import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" //enviar e receber mensagem do canal

	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

	}
}
func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
