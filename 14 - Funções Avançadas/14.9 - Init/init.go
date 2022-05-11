package main

import "fmt"

var n int

// FUNÇÃO INIT INICIALIZA SEMPRE EM PRIMEIRO
func init() {
	fmt.Println("Execuntando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}