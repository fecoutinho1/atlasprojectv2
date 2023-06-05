package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "http://www.bancocn.com/cat.php"

	parameterFile := "parameters.txt"
	parameters, err := loadParameters(parameterFile)
	if err != nil {
		fmt.Printf("Erro ao carregar os parâmetros: %s\n", err.Error())
		return
	}

	wordlistFile := "parameters.txt"
	wordlist, err := loadWordlist(wordlistFile)
	if err != nil {
		fmt.Printf("Erro ao carregar a wordlist: %s\n", err.Error())
		return
	}

	for _, parameter := range parameters {
		for _, payload := range wordlist {
			fullURL := fmt.Sprintf("%s?%s=%s", url, parameter, payload)

			resp, err := http.Get(fullURL)
			if err != nil {
				fmt.Printf("Erro ao acessar a URL: %s\n", err.Error())
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				fmt.Printf("O site aceita o parâmetro '%s' com o valor '%s'\n", parameter, payload)
			} else {
				fmt.Printf("O site não aceita o parâmetro '%s' com o valor '%s'\n", parameter, payload)
			}
		}
	}
}

func loadParameters(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var parameters []string
	for scanner.Scan() {
		parameters = append(parameters, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return parameters, nil
}

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
