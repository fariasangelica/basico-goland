package main

// O import é o que roda. Importamos o "fmt", pois ele tem a função de println. Ou seja, ele printa a tela.
import (
	"fmt"
	ange "strings" // Podemos dar nomes diferentes para os pacotes.
)

// Pacote é um conjunto de funções.
// Existe a possibiilidade de importar outros pacotes fora do Go.


func main() {
    fmt.Println("Hello, world")

	ange.Split("steph","") //mudança do nome do pacote "strings"

}

