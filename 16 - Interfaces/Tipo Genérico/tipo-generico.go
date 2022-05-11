package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(123456))

	// MODO NÃO RECOMENDADO
	mapa := map[interface{}]interface{}{   
		1: 		"String",
		float32(100): true,
		"String":	  "String",
		true:		  float64(12),	
	}

	fmt.Println(mapa)
}