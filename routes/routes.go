package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maickmachado/go-rest-api/controllers"
)

//informa a URL que deverá executar determinada ação
//responsável por indicar quem será responsável quando chegar alguma requisição
func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//quando chegar uma requisição em barra quem vai atender será o controle Home.
	myRouter.HandleFunc("/", controllers.Home)
	//quando chegar uma requisição em barra quem vai atender será o controle ReturnAllArticles.
	myRouter.HandleFunc("/articles", controllers.ReturnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", controllers.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{Id}", controllers.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{Id}", controllers.UpdateArticle).Methods("PUT")
	//utiliza a variável ID para acessar um único artigo
	myRouter.HandleFunc("/articles/{Id}", controllers.ReturnSingleArticle).Methods("GET")
	//significa que qualquer problema que tenha quando subir o servidor ele vai nos apontar
	//ListenAndServe especifica a porta que rodará a aplicação e como lidar com o servidor
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
