package main

import (
	"fmt"
	"testando/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Imprimindo do main!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("email@dominio.com")
	erro2 := checkmail.ValidateFormat("email")
	fmt.Println(erro)
	fmt.Println(erro2)
}
