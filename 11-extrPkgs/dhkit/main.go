package main

import (
	"encoding/json"
	"fmt"

	"github.com/ncostamagna/go-http-utils/meta"
	"github.com/ncostamagna/go-http-utils/response"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mail      string `json:"mail"`
	Age       int    `json:"age"`
}

func main() {
	u := &User{
		FirstName: "Nahuel",
		LastName:  "Costamagna",
		Mail:      "nlcostamagna@gmail.com",
		Age:       32,
	}

	print(response.OK("", u, nil))

	print(response.OK("Pollo con crema jajajaja no se we es un response de prueba", u, nil))

	// 15 es el limite de paginacion por defecto en caso de no asignarle valor
	m, err := meta.New(1, 30, 100, "15")
	if err != nil {
		fmt.Println(err)
	}

	print(response.OK("Se mando un OK aca muy chido", u, m))

	print(response.Created("se mando un Created pero de los que crean bien perron", u, nil))
	print(response.Accepted("se mando un Accepted ¿porque? no se, pero ya esta aceptado", u, nil))

	print(response.BadRequest("Se mando un BadRequest por alguna razon, no se porque, pero ahi esta"))
	print(response.NotFound("Se mando un NotFound, en el bosque de la china la chinita se perdio "))
	print(response.InternalServerError("se mando un InternalServerError, porque se callo el servidor por andar corriendo en la oficina"))

}

func print(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
	fmt.Println()
}
