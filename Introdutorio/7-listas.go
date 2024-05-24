package main

import "fmt"

func main() {

	// Array
	var array [2]string
	array[0] = "Hello"
	array[1] = "world"
	// fmt.Println(array[0], array[1])
	// fmt.Println(array)

    // //Lista
	// numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(numPrimos) 
	// fmt.Println(numPrimos[1]) 
	// fmt.Println(numPrimos[4])
	// fmt.Println(numPrimos[0:3]) //pegando da posição 0 até a 3. Ele não inclui o 3. 
	// fmt.Println(numPrimos[:1]) // tudo antes da posição 1.


	// Slices
	//var slice []string
	slice := make([]string, 5) //defindo um tamanho 5. posições anteriores vazias. Make é porque está criando do zero.
	slice[0] = "Hello"
	slice [1] = "world"
	// fmt.Println(slice[0], slice[1])
	// fmt.Println(slice)
	// fmt.Println(slice[1])
	// fmt.Println(slice[2]) // atribuição a posição 2 do slice
	//slice[2] = "angel"
	// fmt.Println(slice)

	// fmt.Println(len(slice))
	// fmt.Println(slice[2])
	// fmt.Println(slice[3])
	// fmt.Println(slice[5]) //pode criar o slice com qualquer tamanho, mas no nosso caso foi 5.

	// criando com valores - slice
	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	// adicionando mais um número do mesmo tipo na lista. Função append.
	numPares = append(numPares, 14)
	fmt.Println(numPares)

}


// LISTAS

// 1 - Arrays e Slices: Homogêneos
// todos os elementos tem o mesmo tipo [1, 2, 3, 4, 5, 6]
// ["angel", "bento", "goand"]


// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outros
// map[string]int
// {"angel": 28, "bento": 4}
// map[string]string
// {"angel": "farias", "bento": "farias"}

/* Array
- Tamanho fixo, de zero ou mais elementos do mesmo tipo
- Acessamos os valores com índice: a[0], a[1]...
- Função embtida len() retorna o tamanho do array
- Por conta do tamanho fixo, não é tãp usado. Só em casos específicos. */


/* Slice

- Tipo o array, mas sem tamanho fixo.
- Acessamos os valores com índice: a[0], a[1]...
- Função embutida len() retorna o tamanho do slice.
- Função append() usada para adicionar valores.*/