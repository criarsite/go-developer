package main

import (
	"fmt"
	_ "modulo/fundamentos"

	"modulo/estruturadecontrole"
	"modulo/projetos"
)

func main() {

	//fundamentos.Fundamentos()

	var temperaturaDeEbulicacao = projetos.Conversor()
	fmt.Println(temperaturaDeEbulicacao)

	estruturadecontrole.Estruturas()

}
