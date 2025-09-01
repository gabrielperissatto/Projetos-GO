package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>><")
	fmt.Println("GERENCIADOR DE ESTOQUE")
	fmt.Println("PARA CADA OPÇÃO ESCREVA O NÚMERO INDICADO")
	fmt.Println("ADICIONAR PRODUTO [ 1 ] ")
	fmt.Println("BUSCAR PRODUTO   [ 2 ] ")
	fmt.Println("EXIBIR ESTOQUE    [ 3 ]")
	fmt.Println("SAIR              [ 4 ]")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<>>>>>>>>>>>>>><<<<<<<<<<<<<>>>>>>>>>>>>>>><<<<<>>>>>>>")
	for {
		var selection int
		fmt.Print("\nEscolha uma opção: ")
		fmt.Scanln(&selection)
		fmt.Println("OPÇÃO SELECIONADA", "[", selection, "]")

		if selection == 1 {
			adiciona_produto()
		} else if selection == 2 {
			busca_produto()
		} else if selection == 5 {
			remove_produto()
		} else if selection == 3 {
			exibir_estoque()
		} else if selection == 4 {
			fmt.Println("Saindo do programa...")
			break
		} else {
			fmt.Println("ERRO --> ESSA SELEÇÃO NÃO EXISTE. Tente novamente.")
		}
	}
}

// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================

func adiciona_produto() {
	fmt.Println("DIGITE OS DADOS DO PRODUTO A SER ADICIONADO NO ESTOQUE ")

	fmt.Print("NOME DO PRODUTO: ")
	var produto string
	fmt.Scanln(&produto)

	fmt.Print("PREÇO DO PRODUTO: ")
	var preco string
	fmt.Scanln(&preco)

	arquivo := "estoque.txt"
	arqv_mod, err := os.OpenFile(arquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("ERRO AO ABRIR/CRIAR UM ARQUIVO:", err)
		return
	}
	defer arqv_mod.Close()
	produto_estoque := produto + "|" + preco + "\n"
	_, err = arqv_mod.WriteString(produto_estoque)
	if err != nil {
		fmt.Println("ERRO AO ADICIONAR PRODUTO:", err)
		return
	}
	fmt.Println("PRODUTO ADICIONADO COM SUCESSO!")
}

// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================

func exibir_estoque() {
	fmt.Println("\\\\\\\\\\\\\\ESTOQUE////////////")
	arquivo := "estoque.txt"
	arquivo_ex, err := os.ReadFile(arquivo)
	if err != nil {
		fmt.Println("ERRO AO EXIBIR O ESTOQUE !!!!!!!!", err)
		return
	}
	arquivo_string := string(arquivo_ex)
	linhas := strings.Split(arquivo_string, "\n")

	for _, linha := range linhas {

		if len(linha) == 0 {
			continue
		}
		dados := strings.Split(linha, "|")

		if len(dados) == 2 {
			fmt.Println("NOME DO PRODUTO: ", dados[0])
			fmt.Println("PRECO DO PRODUTO: ", dados[1])
			fmt.Println("___________________________________")
		} else {
			fmt.Println("Atenção: Formato de dados inválido na linha:", linha)
		}
	}
}

// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
func busca_produto() {
	fmt.Println("REGISTRO A PROCURAR")
	var chave string
	fmt.Scanln(&chave)

	arquivo := "estoque.txt"
	arquivo_ex, err := os.ReadFile(arquivo)
	if err != nil {
		fmt.Println("ERRO AO BUSCAR ARQUIVO:", err)
		return
	}

	dado := string(arquivo_ex)
	linhas := strings.Split(dado, "\n")

	for _, linha := range linhas {
		if len(linha) == 0 {
			continue
		}

		dados := strings.Split(linha, "|")

		if len(dados) > 0 && dados[0] == chave {

			fmt.Println("ITEM ENCONTRADO! Detalhes do registro:")
			fmt.Println(linha)
			return
		}
	}
	fmt.Println("ITEM NÃO ENCONTRADO.")
}

// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================
// ===================================================================================================================================================================================

func remove_produto() {
	fmt.Println("REGISTRO A REMOVER")
	var chave string
	fmt.Scanln(&chave)

	arquivo := "estoque.txt"
	conteudo, err := os.ReadFile(arquivo)

	if err != nil {
		fmt.Println("ERRO AO BUSCAR ARQUIVO:", err)
		return
	}

	linhas := strings.Split(string(conteudo), "\n")
	var novoConteudo string
	removido := false

	for _, linha := range linhas {
		if strings.HasPrefix(linha, chave+"|") {
			removido = true
			continue
		}
		if len(linha) > 0 {
			novoConteudo += linha + "\n"
		}
	}

	if removido {
		err := os.WriteFile(arquivo, []byte(novoConteudo), 0666)
		if err != nil {
			fmt.Println("ERRO AO REESCREVER O ARQUIVO:", err)
			return
		}
		fmt.Println("REGISTRO REMOVIDO COM SUCESSO!")
	} else {
		fmt.Println("REGISTRO NÃO ENCONTRADO.")
	}
}
