package service

import (
	"package/src/config"
	"package/src/models"
)

func GetById(id string) models.LivroOut {
	db, err := config.OpenConec()
	CheckError(err)

	sql := `SELECT * FROM livros WHERE id=$1`
	rows := db.QueryRow(sql, id)
	CheckError(err)

	var l models.LivroOut

	err = rows.Scan(&l.Id, &l.Titulo, &l.Autor, &l.Genero, &l.Descricao, &l.Preco)
	CheckError(err)

	return l
}
