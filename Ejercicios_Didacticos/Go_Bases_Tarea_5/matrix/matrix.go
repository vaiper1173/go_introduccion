package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	M          [][]float64
	H          int64
	W          int64
	Cuadratica bool
}

func New(v ...[]float64) (*Matrix, error) {
	m := Matrix{
		H: int64(len(v)),
	}

	for i := 0; i < len(v); i++ {
		if i == 0 {
			m.W = int64(len(v[i]))
		} else {
			if m.W != int64(len(v[i])) {
				return nil, errors.New("Error en el tamaño de la matriz, ponga mas atencion Joven")
			}
		}
	}
	m.M = v
	m.Cuadratica = m.H == m.W
	return &m, nil
}

func (m *Matrix) Print() {
	for i := 0; i < len(m.M); i++ {
		fmt.Print("[ ")
		for j := 0; j < len(m.M[i]); j++ {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print((m.M[i][j]))
		}
		fmt.Print(" ]")
		fmt.Println()
	}

	fmt.Println(m.H, "x", m.W)
	fmt.Println("Cuadratica: ", m.Cuadratica)
	fmt.Println()
}
