package atividade6

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
		fmt.Println(HttpGet(site))
	}

}

func HttpGet(site string) (*http.Response, error) {
	resp, err := http.Get(site)
	return resp, err

}
