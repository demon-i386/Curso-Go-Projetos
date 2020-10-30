package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string

	fmt.Print("Digite o primeiro numero: ")
	fmt.Scanf("%s", &text)

	// var num int
	// num, _ = strconv.Atoi(text)
	num, _ := strconv.ParseInt(text, 10, 64)

	fmt.Printf("Numero Convertido: %v\n", num)
}
