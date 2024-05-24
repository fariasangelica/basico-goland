package main

import "fmt"

/* Structs
- Forma de criar a própria estrutura de dados.
- Personalizar de acordo com a sua necessidade.
- Podemos usar vários tipos diferentes. */

// Sintaxe da struct
// type <nome da estrutura> { <campos>}
type Pessoa struct {
	Nome string
	Idade int
}

// prints da estrutura struct

func main() {
	fmt.Println(Pessoa{"Angel", 41})
	fmt.Println(Pessoa{Nome: "Bento", Idade: 4}) // mais aconselhado
	fmt.Println(Pessoa{Nome: "Angel"})

	// variável p1 do tipo pessoa
	p1 := Pessoa{Nome: "Bob", Idade: 2} // struct pessoa atrelada a uma variável p1.
	// fmt.Println(p1.Nome)
	// fmt.Println(p1.Idade)


    // mudar valor da struct
	// p1.Idade = 3
	// fmt.Println(p1.Idade)

	// append com mesmo tipo
	p2 := Pessoa{Nome: "Patrick", Idade: 2}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	//fmt.Println(pessoas)

	 //structs + map
	// alunos := map[string][]Pessoa{}
	// alunos["programação"] = pessoas
	// fmt.Println(alunos)

	// var alunos = map[string] []Pessoa{
	// 	"Programação": {{Nome: "Angel", Idade: 39}, {Nome: "Bento", Idade: 4}},
	// 	"Engenharia": {{Nome: "Angel2", Idade: 21}, {Nome: "Bento2", Idade: 2}},

	// }
    // fmt.Println(alunos)


	// struct herdando campos de outras struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}

type Profissao struct {
	Pessoa
	Tipo string
}