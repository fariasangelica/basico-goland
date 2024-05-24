// Retorno negativo

/* Nesta tarefa simples, você recebe um número e deve torná-lo negativo. Mas talvez o número já seja negativo.*/

// 1 - recebe um numero.
// 2 - verificar se o numero é negativo, se for, já retorna.
// 3 - se não for, transofrma-lo em negativo.
// 4 - 5 = (-1)

package kata

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	}
	negativeNumer := x * (-1)

	return negativeNumer
}
