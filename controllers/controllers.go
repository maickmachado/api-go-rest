package controllers

import (
	"fmt"
	"net/http"
)

//toda vez que chegar uma requisição em Home(), vai responder exibindo essa mensagem "Home Page"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
