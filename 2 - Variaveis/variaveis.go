package main

import "fmt"

func main() {
	var variavel1 string = "var 1"
	variavel2 := "var 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		var3 string = "123"
		var4 string = "123"
	)

	fmt.Println(var3, var4)

	var5, var6 := "var5", "var6"

	const const1 string = "const 1"
	fmt.Println(const1)

	fmt.Println(var5, var6)

	var5, var6 = var6, var5

	fmt.Println(var5, var6)
	//o go nao deixa vc criar variavel e n usar
	// o go nao deixa vc importar e n usar
}
