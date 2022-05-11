// STRUCTS EM GO É O MESMO QUE USAR CLASSES EM LINGUAGEM ORIENTADAS A OBJETOS, 
//POIS EM GO NÃO TEM CLASSES, (STRUCTS É COLEÇÃO DE CAMPOS)
package main

import "fmt"

type usuario struct {
	nome 		string
	idade 		uint8
	endereco 	endereco
}

type endereco struct {   
	logradouro string
	numero 	   uint8
}

func main()  {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Matheus"
	u.idade = 4
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bois", 0}

	// FORMA MAIS SIMPLES DE DECLARAR STRUCT
	usuario2 := usuario{"Matheus", 4, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Matheus"}
	fmt.Println(usuario3)
}