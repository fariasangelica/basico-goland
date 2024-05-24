package main

import "fmt"

func main() {
	fmt.Println(soma(42,13))

	sub := subtracao(10,5)
	 fmt.Println(sub)

	nome := printaNome("Angel")
	fmt.Println(nome)
}

func printaNome(nome string) string{
	return nome
}

func soma(x int, y int) int {
	return x + y // o que a função retorna.
}

// Função começando com letra minúscula:
// Função é PPRIVADA.
// Só pode ser utilizada no próprio pacote.
    
// Função criada nova criada para somar dois números.
func subtracao(x int, y int) int {
	return x - y // o que a função retorna.
}