package main

import (
	"fmt"
)


type Pessoa struct {
	nome string
	idade int
}

func main() {
	var nomeInput string
	var idadeInput int

	fmt.Println("Digite o nome: ")
	fmt.Scan(&nomeInput)

	fmt.Println("Digite a idade: ")
	fmt.Scan(&idadeInput)

	pessoa := Pessoa{
		nome: nomeInput,
		idade: idadeInput,
	}

	fmt.Println("")
	
	fmt.Println("Nome: ", pessoa.nome)
	fmt.Println("Idade: ", pessoa.idade)
}