package main

//goroutine funtion capaz de executar de modo concorrente com outras funcoes
import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
	//
}
func main() {
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)
}
