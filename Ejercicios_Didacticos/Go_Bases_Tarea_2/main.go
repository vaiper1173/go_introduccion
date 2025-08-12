package main

import "fmt"

func main() {

	array := [5]int{4, 2, 5, 6, 7}

	// realizar la funcionalidad
	for i := 0; i < 20; i++ {
		array[0] = array[0] + 1
		array[1] = array[1] + 1
		array[2] = array[2] + 1
		array[3] = array[3] + 1
		array[4] = array[4] + 1
	}

	fmt.Println("Los valores del array son: ", array)
}
