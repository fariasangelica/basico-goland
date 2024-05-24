package main

import "fmt"

// Estrutura maps
func main() {
	// idade := map[string]int{}
	// idade["angel"] = 39
	// idade["bento"] = 4
	// fmt.Println(idade)
	// fmt.Println(idade["angel"])
	// fmt.Println(idade["bento"])

	// Outra forma da estrutura maps
	anoNasc := map[string]int{
		"angel": 1985,
		"bento": 2008,

	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["angel"])
	fmt.Println(anoNasc["bento"])

	// Adicionar novos elementos no array - maps
	anoNasc["golandDoZero"] = 2024
    fmt.Println(anoNasc)
}

// Podemos misturar tipos em Maps