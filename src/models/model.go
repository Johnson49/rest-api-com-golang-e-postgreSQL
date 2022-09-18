package models


type Livro struct{
	Titulo string
	Autor string
	Genero string
	Descricao string
	Preco float64
}

type LivroOut struct{
	Id string `json:"id"`
	Titulo string `json:"título"`
	Autor string	`json:"autor"`
	Genero string	`json:"gênero"`
	Descricao string `json:"descrição"`
	Preco float64 `json:"preço"`
}