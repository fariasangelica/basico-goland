package main

import (
	"fmt"
	
)

// Loops
// Laços de repetição
// Repetir tarefas

func main() {
	 
	// sum := 0

	// for i := 0; i < 10; i++ { // O "i" indica a posição inicial. O "i" começa em zero. Loop acontece até que o numero seja menor que 10.
	// 	fmt.Println("Sum:", sum)
	// 	fmt.Println("Indice:", i)
	//     sum += i  //sum = sum + i           // i ++ => soma 1  // i -- => subtraindo 1

		// final do loop é adicionado 1 ao indice i
		// é a mesma coisa que: sum = sum + i
		// sum -= -> sum = sum - i
	// }
	// fmt.Println(sum)



	// // Loop infinito
	// for {
	// 	fmt.Println("Goland do zero")
	// 	time.Sleep(2 * time.Second)
	// }


	// for range
	frutas :=[]string{"laranja", "maça", "banana", "uva", "kiwi"}
	for _, fruta := range frutas { // underscore é a posição.
		fmt.Println(fruta)
	}
}