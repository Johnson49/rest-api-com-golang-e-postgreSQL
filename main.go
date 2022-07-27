package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"

)

type Livro struct {
	Id        string     `json: "id"`
	Titulo    string  `json: "titulo"`
	Autor     string  `json: "autor"`
	Genero    string  `json: "genero"`
	Descricao string  `json: "descricao"`
	Preco     float32 `json: "preco"`
}


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/biblioteca", pegarLivros).Methods("GET")
	router.HandleFunc("/biblioteca/{id}", pegarLivro).Methods("GET")
	router.HandleFunc("/biblioteca/add", adicionarLivro).Methods("POST")
	router.HandleFunc("/biblioteca/{id}", excluirLivro).Methods("DELETE")
	router.HandleFunc("/biblioteca/{id}", atualizarLivro).Methods("PUT")
  
	const port = ":8000"
	fmt.Println("Server is listenning on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func pegarLivros(w http.ResponseWriter, r *http.Request) {
	livros, err := os.ReadFile("./database.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write(livros)
}

func pegarLivro(w http.ResponseWriter, r *http.Request) {
  
	dados, err := os.ReadFile("./database.json")
	if err != nil {
		log.Fatal(err.Error())
	}
  
  var livros []Livro
  _ = json.Unmarshal(dados, &livros)
  
  params := mux.Vars(r)
    for _, item := range livros {
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }

   json.NewEncoder(w).Encode("Livro não encontrado.")

}

func adicionarLivro(w http.ResponseWriter, r *http.Request) {
  var pacote []Livro
  
	livros, _ := os.ReadFile("./database.json")

  _ = json.Unmarshal(livros, &pacote)

	var livro Livro
	
    _ = json.NewDecoder(r.Body).Decode(&livro)

  pacote = append(pacote, livro)

  dados, erro := json.Marshal(pacote)
  
	if erro != nil {
	panic(erro)
	}
  
  err := os.WriteFile("./database.json", dados, 0755)

  if err != nil{
    panic(err.Error())
  }
  json.NewEncoder(w).Encode("Livro adicionado com sucessor.")
}


func excluirLivro(w http.ResponseWriter, r *http.Request){
  var pacote_de_livros []Livro
  
	livros, _ := os.ReadFile("./database.json")

  err := json.Unmarshal(livros, &pacote_de_livros)

  if err != nil{
    panic(err.Error())
  }

  params := mux.Vars(r)
    for index, item := range pacote_de_livros {
        if item.Id == params["id"] {
            pacote_de_livros = append(pacote_de_livros[:index], pacote_de_livros[index+1:]...)
          
  dados, erro := json.Marshal(pacote_de_livros)
  
	if erro != nil {
	panic(erro)
	}
  
  err := os.WriteFile("./database.json", dados, 0755)

  if err != nil{
    panic(err.Error())
  }
  
             json.NewEncoder(w).Encode("Livro excluído com sucessor.")
            
        }
     
    }
}

func atualizarLivro(w http.ResponseWriter, r *http.Request){

   var pacote_de_livros []Livro
  
	livros, _ := os.ReadFile("./database.json")

  err := json.Unmarshal(livros, &pacote_de_livros)

  if err != nil{
    panic(err.Error())
  }

  params := mux.Vars(r)
    for index, item := range pacote_de_livros {
        if item.Id == params["id"] {
            pacote_de_livros = append(pacote_de_livros[:index], pacote_de_livros[index+1:]...)
        }
    }
	var livro Livro
	
    _ = json.NewDecoder(r.Body).Decode(&livro)

  pacote_de_livros = append(pacote_de_livros, livro)

      
  dados, erro := json.Marshal(pacote_de_livros)
  
	if erro != nil {
	panic(erro)
	}
  
  resul := os.WriteFile("./database.json", dados, 0755)

  if resul != nil{
    panic(err.Error())
  }
  
             json.NewEncoder(w).Encode("Livro atualizado com sucessor.")
            

          
}