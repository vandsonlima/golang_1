package main

import "fmt"

func main() {
	// inteiros e reais

	/*inteiros
	    int8
		int16
		int32
		int64
	*/
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 int = 100000000 // de acordo com a arq do computador
	fmt.Println(numero2)

	numero3 := 123123123 // inferencia de tipo de acordo com a arq do computador
	fmt.Println(numero3)

	//unsigned int -> uint

	//alias para tipos de dados

	var numero4 int32 = 1234
	var numero5 rune = 1234

	fmt.Println(numero4, numero5) // int32 = rune

	// byte = int8

	var numeroReal float32 = 123.45
	var numeroReal2 float64 = 123.45
	numeroReal3 := 123.45 // inferencia para float64

	fmt.Println(numeroReal, numeroReal2, numeroReal3)

	//FIM NUMEROS REAIS

	//STRINGS
	var str string = "texto"
	fmt.Println(str)

	str2 := "texto 2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)

	//teste

}
