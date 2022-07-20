package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maickmachado/go-rest-api/models"
)

//toda vez que chegar uma requisição em Home(), vai responder exibindo essa mensagem "Home Page"
func Home(w http.ResponseWriter, r *http.Request) {
	//imprime na URL w
	fmt.Fprint(w, "Home Page")
	//imprime no console
	fmt.Println("Endpoint hit: home page")
}

//função que serve para colocar dados dentro de Articles
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return all articles")
	//codifica os arrays dos artigos em um JSON string e então coloca eles no models.Articles
	json.NewEncoder(w).Encode(models.Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return single article")
	//retorna as variáveis da requisição atual é um map[string]string
	vars := mux.Vars(r)
	key := vars["Id"]
	//esse fmt pega aceita um writer e um formato definido
	//no caso ela pega o w (URL recebida) e escreve no formato key: 1
	fmt.Fprintf(w, "key: %v\n", key)
	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, value := range models.Articles {
		if value.Id == key {
			//codifica os arrays dos artigos em um JSON string e então coloca eles no value (models.Articles)
			json.NewEncoder(w).Encode(value)
		}
	}

}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return a create article")
	//lê toda a informação até ocorrer um erro e retorna um []byte e um erro
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	// unmarshal this into a new Article struct
	json.Unmarshal(reqBody, &article)
	// append this to our Articles array.
	models.Articles = append(models.Articles, article)
	//codifica os arrays dos artigos em um JSON string e então coloca eles no article (models.Articles)
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return a delete article")
	vars := mux.Vars(r)
	id := vars["Id"]
	for index, value := range models.Articles {
		if value.Id == id {
			// updates our Articles array to remove the article
			models.Articles = append(models.Articles[:index], models.Articles[index+1:]...)
		}
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return a update article")
	vars := mux.Vars(r)
	id := vars["Id"]
	//lê toda a informação até ocorrer um erro e retorna um []byte e um erro
	// get the body of our PUT request
	var updatedEvent models.Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	// unmarshal this into a new Article struct
	json.Unmarshal(reqBody, &updatedEvent)
	for i, value := range models.Articles {
		if value.Id == id {
			value.Title = updatedEvent.Title
			value.Desc = updatedEvent.Desc
			value.Content = updatedEvent.Content
			models.Articles[i] = value
			//codifica os arrays dos artigos em um JSON string e então coloca eles no value (models.Articles)
			json.NewEncoder(w).Encode(value)
		}
	}
}
