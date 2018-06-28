package main

import (
	"encoding/json"
	"fmt"

	"github.com/fyukiobr/workshop-go/domain"
)

func main() {
	relatives := map[string]interface{}{}
	relatives["father"] = "Luigi"
	relatives["mother"] = "Rosana"
	relatives["brothers"] = []string{"Fernanda", "Gian"}

	usuario := domain.User{
		Name:      "Nome teste",
		Age:       15,
		Phones:    []string{"123", "5656", "2345667454"},
		Relatives: relatives,
	}

	//fmt.Printf("User: %s\n", usuario)
	fmt.Printf("User: %v\n", usuario)
	fmt.Printf("User: %+v\n", usuario)

	for i, v := range usuario.Phones {
		fmt.Printf("Fone[%d]: %s\n", i, v)
	}

	for i, v := range usuario.Relatives {
		fmt.Printf("Parente[%s]: %s\n", i, v)
	}

	nome, encontrado := usuario.Relatives["father"]
	if encontrado {
		fmt.Printf("Father: %s\n", nome)
	}

	jsonret, err := json.Marshal(usuario)
	if err == nil {
		fmt.Println(string(jsonret))
	}

}
