package service

import (
	"github.com/Johnson49/rest-api-com-golang-e-postgreSQL/src/config"
	"github.com/Johnson49/rest-api-com-golang-e-postgreSQL/src/models"
)

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}


func GetAll() ([]models.LivroOut){
	db, err := config.OpenConec()
	if err != nil {
		panic(err)
	}
	var livros []models.LivroOut


	rows, err := db.Query("SELECT * FROM livros")
    CheckError(err)

	defer rows.Close()
	
	for rows.Next(){
		var l models.LivroOut

		err := rows.Scan(&l.Id, &l.Titulo, &l.Autor, &l.Genero, &l.Descricao, &l.Preco,)
		CheckError(err)

		livros = append(livros, l)
	}
	
	return livros

}