// GO NÃO TEM  HERANÇA ESSE É SOMENTE UM EXEMPLO, 
//SOBRE O QUE A LINGUAGEM CHEGA MAIS PROXIMO DE HERANÇA
package main

import "fmt"

type pessoa struct {
	nome 		string
	sobrenome 	string
	idade 		uint8
	altura 		uint8

}

type estudante struct {
	// HERANÇA EM GO BASICAMENTE é PEGAR O NOME DE OUTRA STRUCT
	pessoa
	curso		string
	faculdade 	string

}

func main()  {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.altura)
	
}