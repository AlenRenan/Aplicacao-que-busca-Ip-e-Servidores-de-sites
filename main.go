/*
	A aplicação irá bucar em um endereço web o ip publico
	do endereço e o nome do servidor hospedado
*/
package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	aplicacao := app.Gerar()       //Chama a função Gerar da classe app
	erro := aplicacao.Run(os.Args) // Parametro que fara com que os comandos dos S.O sejam reconhecidos. Como retorna um erro, tem que ter uma verificação

	if erro != nil {
		log.Fatal(erro) //gera um log de erro e para a execução do programa
	}

}
