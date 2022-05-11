// ESTRUTURAS DE CONTROLE (IF ELSE)

package main

import "fmt"

func main()  {
	fmt.Println("Estruturas de controle")

	numero := -5

	// O IF VAI AVALIAR UMA CONDIÇÃO SE É VERDADEIRA
	if numero > 15 {
		fmt.Println("Maior que 15")
		// SE NÃO FOR VERDADEIRA ELA CAI NO ELSE
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
		// ELSE IF AVALIA UMA OUTRA CONDIÇÃO
	} else if numero < - 10 {
		fmt.Println("Número é menor que -10 ")
	} else {
		fmt.Println("Entre 0 e -10")
	}
	
}