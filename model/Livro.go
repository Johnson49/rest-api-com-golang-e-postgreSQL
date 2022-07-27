package model

type Livro struct {
	Id        string     `json: "id"`
	Titulo    string  `json: "titulo"`
	Autor     string  `json: "autor"`
	Genero    string  `json: "genero"`
	Descricao string  `json: "descricao"`
	Preco     float32 `json: "preco"`
}