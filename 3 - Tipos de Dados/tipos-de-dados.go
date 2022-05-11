package main

import ( 
	"fmt"
	"errors"
)
		

func main()  {
	// int é a quantidade de bit que ocupa 
	//tipos de números inteiros no Go ---> int8, int16, int32, int64 <-------
	// "int" sozinho usa a arquitetura do computador que esta sendo executada x64 ou x32 
	
	//var numero int16 = 100
	//fmt.Println(numero)

	numero := -100000000000
	fmt.Println(numero)

	// "uint" unsygned int (int sem sinal)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//"alias" (apelido)  
	//"rune" é exatamente igual o int32"
	var numero3 rune = 12456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// números reais
	// float32 e float64
	var numeroReal1 float32 = 123.45 // sempre tem que ser ponto a vírgula não vai funcionar
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000000.45
	fmt.Println(numeroReal2)

	//pega o valor da arquitetura do computador 64 ou 32
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NÚMEROS REAIS

	// tipos parea declarar uma string
	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// faz referẽncia a número inteiro
	char := 'A'
	fmt.Println(char)

	// FIM STRINGS

	// VALOR VAZIO PARA (string) e VALOR ZERO PARA (número)
	var texto int16
	fmt.Println(texto)

	//NÃO USADO
	//texto := 5
	//fmt.Println(texto)

	// valores booleanos true e false
	var booleano1 bool
	fmt.Println(booleano1)

	// valor erro igual a nil (erro nome da variável), (error é o tipo da variável)
	// (errors é o nome do pacote)
	// valor tipo erro não é string é erro mesmo
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}