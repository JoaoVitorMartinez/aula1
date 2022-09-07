package atividade6

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func Monitoramento() {
	sites := []string{"https://www.google.com"}

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arquivo)
	}
	leitor := bufio.NewReader(arquivo)
loop:
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break loop
		}

	}
	arquivo.Close()

	for i, site := range sites {
		fmt.Println(i, site)
		arquivo, err = os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println(err)
		}
		status, _ := HttpGet(site)
		fmt.Println(status)

		if status.StatusCode == 200 {
			current := time.Now().Format("02/01/2006 15:04:05 UTC - ")
			arquivo.WriteString(current + site + " Status: OK\n")
		} else {
			current := time.Now().Format("02/01/2006 15:04:05 UTC - ")
			arquivo.WriteString(current + site + "Status: Error\n")
		}

		arquivo.Close()

	}

	fmt.Println()
	fmt.Println("Exibindo conteúdo dos logs")
	fmt.Println()
	log, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(log))

}

func HttpGet(site string) (*http.Response, error) {
	resp, err := http.Get(site)
	return resp, err

}
