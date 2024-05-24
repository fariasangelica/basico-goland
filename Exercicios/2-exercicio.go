
// Escreva uma função que aceite um número inteiro e uma string como parâmetros
//e retorne uma string de s repetida exatamente n vezes.

//Exemplos (input -> output)

//6, "I"     -> "IIIIII"
//5,  "Hello"  -> "HelloHelloHelloHelloHelloHello"

package kata

func RepeatStr(repetitions int, value string) string {
    var repeatString string
	for i:=0;i<repetitions;i++{
		repeatString = repeatString + value
        
	}
	return repeatString
}