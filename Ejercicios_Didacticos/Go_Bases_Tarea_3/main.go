package main

import (
	"fmt"
)

func main() {
	var numeros []int // Slice para almacenar los números en teoria
	var valor int

	fmt.Println("Ingrese números (0 para finalizar):")

	for {
		fmt.Print("Número: ")
		fmt.Scan(&valor)

		if valor == 0 {
			break // Salimos del bucle
		}

		numeros = append(numeros, valor) // Agregamos al slice en la practica
	}

	fmt.Println("\nValores ingresados:")
	fmt.Println(numeros) //si funciono XD
}
