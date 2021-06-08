package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

/*
	Retona a aplicação de linha de comando
	pronta para ser executada
*/
func Gerar() *cli.App { //G maiusculo para ser público

	/*
		.Name, .Usage, .UsageText, .Version são da biblioteca cli
	*/
	app := cli.NewApp()                                                                                                                                                                                                                                                                     //cria uma variável que chama a função
	app.Name = "Aplicação de linha de comando"                                                                                                                                                                                                                                              //Imprime o nome do programa
	app.Usage = "Busca IPs e nomes de servidores na Internet"                                                                                                                                                                                                                               //Imprime para que serve o programa
	app.UsageText = "Para utilizar o programa deve-se : \n\t 1º Para conseguir o ip --> go run (nome do arquivo onde está o main).go ip --host (url do site)\n\t 2º  Para conseguir o nome do servidor --> go run (nome do arquivo onde está o main).go servidores --host (url do site) \n" //Imprime como ensina o programa
	app.Version = "1.0.0 \t June 08, 2021"                                                                                                                                                                                                                                                  //Imprime a versão

	flags := []cli.Flag{ //Comando para executar na linha de comando a partir dos parâmetros
		cli.StringFlag{
			Name:  "host",       //Valor que será enviado para a função buscarIps
			Value: "google.com", //Valor padrão caso não consiga executar o que foi mandado
		},
	}

	/*
		Cria um slice com dois comandos
	*/
	app.Commands = []cli.Command{ //Slice de comandos que a aplicação irá executar
		/*
			Comandos para conseeguir o ip
		*/
		{
			Name:   "ip", //comando que será utilizado na linha de comando
			Usage:  "Busca Ip de endreço web",
			Flags:  flags,     //recebe a variavel externa flags com os comandos criados
			Action: buscarIps, //chama a função buscarIps
		},
		/*
			Comandos para conseeguir o nome do servidor
		*/
		{
			Name:   "servidores", //Comando que será utilizado na linha de comando
			Usage:  "Busca nome dos servidores",
			Flags:  flags,         //recebe a variavel externa flags com os comandos criados
			Action: buscaServidor, //chama a função buscarIps

		},
	}
	return app
}

/*
	Função responsável por buscar os ips de um
	site web
*/
func buscarIps(c *cli.Context) {

	host := c.String("host") //Armazena o valor retornado da flag e armazena na variavel

	/*
		Comando que utiliza a biblioteca "net" que buscará o ip
	*/
	ips, erro := net.LookupIP(host) //Retorna um slice de ips e um erro

	/*
		Caso de erro, ele será imprimido na linha de comando e o programa será finalizado
	*/
	if erro != nil {
		log.Fatal(erro)
	}

	/*
		Loop que rodará de acordo com a quantidade de ips existentes e
		imprimirá os mesmos na linha de comando
	*/
	for _, ips := range ips {

		fmt.Println(ips)

	}

}

/*
	Função responsável por peegar o nome dos serivores
	de um site web
*/
func buscaServidor(c *cli.Context) {

	host := c.String("host")

	servidores, erro := net.LookupNS(host) //Função que retorna o nome do servidor do host e um erro
	/*
		Caso de erro, ele será imprimido na linha de comando e o programa será finalizado
	*/
	if erro != nil {
		log.Fatal(erro)
	}
	/*
		Loop que rodará de acordo com a quantidade de ips existentes e
		imprimirá os mesmos na linha de comando
	*/
	for _, servidores := range servidores {

		fmt.Println(servidores.Host) //.host serve para mostrar apenas o valor da string
	}
}
