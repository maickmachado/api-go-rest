package main

import (
	"fmt"

	"github.com/maickmachado/go-rest-api/models"
	"github.com/maickmachado/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o projeto...")
	models.Articles = []models.Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	routes.HandleRequest()

}
