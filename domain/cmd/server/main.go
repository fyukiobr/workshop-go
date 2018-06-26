package main

import (
	"fmt"

	"github.com/fyukiobr/workshop-go/domain"
)

func main() {
	usuario := domain.User{Name: "Nome teste", Age: 15, Phones: []string{"123", "5656", "2345667454"}}

	fmt.Printf("User: %s\n", usuario)
	fmt.Printf("User: %v\n", usuario)
	fmt.Printf("User: %+v\n", usuario)

	for i, v := range usuario.Phones {
		fmt.Printf("Fone[%d]: %s\n", i, v)
	}
}
