package service

import (
	"package/src/config"
	"package/src/models"
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