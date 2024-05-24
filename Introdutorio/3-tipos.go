package main

import (
	"fmt"
)

func main() {
	//bool (true/false)
	fmt.Print("Type: %T - Value: %v\n", true, true)

	//string - sequência de bytes
	fmt.Print("Type: %T - Value: %v\n", "ange", "ange")
	fmt.Print("Type: %T - Value: %v\n", "1", "1")

	//int - inteiros
	fmt.Print("Type: %T - Value: %v\n", 1, 1)

	//float (float64/float31) - decimal
	fmt.Print("Type: %T - Value: %v\n", 1.233, 1.233)

}

/* Tipos:

bool (true/false)
string - sequência de bytes
int - inteiros
float (float64/float31) - decimal */