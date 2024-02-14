package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste 1"
		variavel4        = 123321
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6, variavel7 := "variavel 5", "varivavel 6", 567
	fmt.Println(variavel5, variavel6, variavel7)

	const constante1 string = "const1"
	fmt.Println(constante1)

	//teste invertendo

	variavel5, variavel6 = variavel6, variavel5 // tem q ser do mesmo tipo

	fmt.Println(variavel5, variavel6)

}
