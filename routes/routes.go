package routes

import (
	"log"
	"net/http"

	"github.com/maickmachado/go-rest-api/controllers"
)

//informa a URL que deverá executar determinada ação
//responsável por indicar quem será responsável quando chegar alguma requisição
func HandleRequest() {
	//quando chegar uma requisição em barra quem vai atender será o controle Home.
	http.HandleFunc("/", controllers.Home)
	//significa que qualquer problema que tenha quando subir o servidor ele vai nos apontar
	//ListenAndServe especifica a porta que rodará a aplicação e como lidar com o servidor
	//nil indica que irá rodar na configuração padrão
	log.Fatal(http.ListenAndServe(":8000", nil))
}
