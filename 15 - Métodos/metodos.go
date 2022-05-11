// MÉTODO FAZ PARTE DE UMA ESTRUTURA

package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

// MÉTODO, VALOR A SER SUBSTITUIDO NO LUGAR DE %S
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
	// \n salva na linha de baixo
}

// METÓDO
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// MÉTODO
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}	
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}