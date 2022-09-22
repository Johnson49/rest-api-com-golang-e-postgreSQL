package controller

import (
	"encoding/json"
	"net/http"
	"github.com/Johnson49/rest-api-com-golang-e-postgreSQL/src/service"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PegarTodosOsLivros(w http.ResponseWriter, _ *http.Request) {
	livros := service.GetAll()
	
	json, err := json.Marshal(livros)
	CheckError(err)

	w.Write(json)
}

func PegarLivroPeloID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	livro := service.GetById(id)
	json, err := json.Marshal(livro)
	CheckError(err)

	w.Write(json)
}

func AtualizarLivro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	resultado := service.Put(id, r.Body)
	if resultado != 1 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Um erro ocoreu ao tentar atualizar livro.")
		return
	}
	json, err := json.Marshal(resultado)
	CheckError(err)

	w.Write(json)
}

func ExcluirLivro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	resultado := service.Delete(id)
	if resultado != 1 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Um erro ocoreu ao tentar excluir livro.")
		return
	}

	json, err := json.Marshal(resultado)
	CheckError(err)

	w.Write(json)
}

func RegistrarLivro(w http.ResponseWriter, r *http.Request) {
	resultado := service.Post(r.Body)

	if resultado != 1 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Um erro ocoreu ao tentar registrar livro.")
		return
	}

	json, err := json.Marshal(resultado)
	CheckError(err)
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}
