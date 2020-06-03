package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//Constantes utilizadas para numerar as opções do fluxo do sistema
const (
	Monitoramento = 1
	ExibirLogs    = 2
	Sair          = 0
)

//Constantes usadas no controle do monitoramento dos sites
const (
	Delay                    = 2
	MonitoramentosExecutados = 2
	PermissaoArquivo         = 0666
)

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		var comando int = leComando()

		switch comando {
		case Monitoramento:
			monitora()
		case ExibirLogs:
			imprimeLogs()
		case Sair:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida, selecione a opção correta...")
		}
	}
}

func exibeIntroducao() {
	versao := 1.0
	fmt.Println("**********************************")
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("**********************************")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("**********************************")
}

func leComando() int {
	var comando int = -1
	/**
	* No Windows é necessário colocar o \n no fim,
	* para que o Scanf espere a entrada de dados
	* Referências:
	* 	https://stackoverflow.com/questions/17401709/why-doesnt-fmt-scanf-in-go-wait-for-user-input
	 */
	fmt.Scanf("%d\n", &comando)
	fmt.Println("O comando seleciona foi o:", comando)
	fmt.Println("**********************************")
	return comando
}

func monitora() {
	fmt.Println("Iniciando monitoramento...")

	sites := leSitesDoArquivo()

	for i := 0; i < MonitoramentosExecutados; i++ {
		for i, site := range sites {
			fmt.Println("Monitorando o site", i+1, ":", site)
			monitoraSite(site)
		}
		fmt.Println("")
		time.Sleep(Delay * time.Second)
	}

}

func monitoraSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if 200 == resp.StatusCode {
		fmt.Println("O site", site, "foi carregado com sucesso :)")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "está com problemas :<")
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	sites := []string{}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if io.EOF == err {
			break
		}

		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, PermissaoArquivo)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString("date:[" + time.Now().Format("02/01/2006 15:04:05") + "] - site:[" + site + "] - online:[" + strconv.FormatBool(status) + "]\n")

	arquivo.Close()

}

func imprimeLogs() {

	fmt.Println("Exibindo logs...")
	bytesArquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(bytesArquivo))
}
