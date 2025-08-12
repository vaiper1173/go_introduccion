package main

import (
	"fmt"
)

const notebook string = "notebook"
const tv string = "tv"
const heladera string = "heladera"
const monitor string = "monitor"
const camara string = "camara"
const pollo string = "pollo en crema joven una delicia"
const notFound string = "No encontrado"

func main() {
	var array []string
	fmt.Println("Seleccione valores, para salir use el cero (0)")

	for {
		var value string
		// al parecer \n es necesario debido a que cuando se da enter el programa interpreta que hay otra cadena igresada como vacia
		fmt.Scanf("%s\n", &value) //"%s" significa "lee una cadena de texto (string) sin espacios" .

		if value == "0" {
			break
		}

		switch value {
		case "10":
			array = append(array, notebook)
		case "15":
			array = append(array, tv)
		case "21":
			array = append(array, heladera)
		case "27":
			array = append(array, monitor)
		case "35":
			array = append(array, camara)
		case "69":
			array = append(array, camara)
		default:
			array = append(array, notFound)
		}

		fmt.Println("Usted metio los valores de: ", array)
	}

}
