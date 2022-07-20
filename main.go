package main

import (
	"fmt"

	"github.com/maickmachado/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o projeto...")
	routes.HandleRequest()
}
