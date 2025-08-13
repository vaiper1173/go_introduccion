package main

import (
	"fmt"

	"github.com/vaiper1173/go_introduccion/Ejercicios_Didacticos/Go_Bases_Tarea_5/matrix"
)

func main() {
	m, err := matrix.New([]float64{2, 7, 15, 10}, []float64{4, 4, 7, 10}, []float64{5, 6, 1, 10}, []float64{10, 10, 10, 10})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()
}
