package main

import "fmt"

func  main()  {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//GO SÓ FUNCIONA COM NÚMERO DO MESMO TIPO NÃO PODE USAR INTS DIFERENTE
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	varriavel2 := "String2"
	fmt.Println(variavel1, varriavel2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2) // MAIOR
	fmt.Println(1 >= 2) // MAIOR IGUAL
	fmt.Println(1 == 2) //IGUAL-IGUAL
	fmt.Println(1 <= 2) // MENOR IGUAL
	fmt.Println(1 < 2) //MENOR
	fmt.Println(1 != 2) //DIFERENTE
	// FIM DOS RELACIONAIS

	// OPERADORES LÓGIDOS
	fmt.Println("-------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // VALOR DIFERENTE 
	fmt.Println(verdadeiro || falso) // UM OU OUTRO (OU)
	fmt.Println(!verdadeiro) // ! VALOR DE NEGAÇÃO, INVERTER O VALOR DA VARIÁVEL
	fmt.Println(!false)
	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // mesma coisa que numero = numero + 15
	// INCREMENTANDO VALORES
	fmt.Println(numero)

	// DECREMENTANDO NÚMERO
	numero--
	numero-=20 // numero = numero - 20

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	// OPERADORES TERNÁRIOS EM GO SÓ EXISTE DESSA FORMA
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
		} else {
			texto = "Menor que 5"	
		}

		fmt.Println(texto)
}