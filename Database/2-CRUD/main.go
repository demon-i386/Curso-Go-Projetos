package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/api/v1" // Ás rotas iram chamar as funções desse package.
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user", api.CreateHTTP).Methods(http.MethodPost)     // Está é a rota que cria um usuário no banco de dados.
	router.HandleFunc("/api/v1/user", api.SearchAllHTTP).Methods(http.MethodGet)   // Está é a rota que puxa todos os usuários do banco de dados.
	router.HandleFunc("/api/v1/user/{id}", api.SearchHTTP).Methods(http.MethodGet) // Está é a rota que puxa os usuários por id do banco de dados.

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
