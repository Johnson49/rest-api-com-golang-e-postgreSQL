package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"package/src/controller"
)

func main() {
	godotenv.Load()
	porta := os.Getenv("PORTA")
	router := mux.NewRouter()

	route := router.PathPrefix("/biblioteca").Subrouter()

	route.HandleFunc("/", controller.PegarTodosOsLivros).Methods("GET")
	route.HandleFunc("/livro", controller.PegarLivroPeloID).Methods("GET")
	route.HandleFunc("/livro", controller.AtualizarLivro).Methods("PUT")
	route.HandleFunc("/livro", controller.ExcluirLivro).Methods("DELETE")
	route.HandleFunc("/registro", controller.RegistrarLivro).Methods("POST")

	fmt.Printf("Server is listenning on http://localhost%s", porta)
	log.Fatal(http.ListenAndServe(porta, router))

}
