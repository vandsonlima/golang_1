package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook.gmail.com")
	// para remover uma dependencia nao usada, basta executar go mod tidy
	fmt.Println(erro)
}
