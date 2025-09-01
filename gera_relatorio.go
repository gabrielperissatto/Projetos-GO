package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := "salarios.csv"

	s, err := os.ReadFile(a)
	if err != nil {
		fmt.Println("ERRO AO ABRIR O ARQUIVO")
		return
	}

	fmt.Println("ARQUIVO ABERTO :3")

	textoDoArquivo := string(s)

	linhas := strings.Split(textoDoArquivo, "\n")

	fmt.Println("\n--- DADOS DOS FUNCIONÁRIOS ---")

	for i, linha := range linhas {

		//if i == 0 || linha == "" {
		//	continue
		//}
		partes := strings.Split(linha, "|")

		if len(partes) == 3 {

			fmt.Printf("FUNCIONÁRIO 🔎", i)
			fmt.Println("")
			fmt.Printf("\tNome 🧑: %s\n", partes[0])
			fmt.Printf("\tCargo >>>: %s\n", partes[1])
			fmt.Printf("\tSalário 💸: %s\n", partes[2])
			fmt.Println("--------------------")
		}
	}
}
