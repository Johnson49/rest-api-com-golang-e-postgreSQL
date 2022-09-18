package service

import (
	"encoding/json"
	"io"
	"package/src/config"
	"package/src/models"
)



func Put(id string, body io.ReadCloser) int64 {
	var l models.Livro
	json.NewDecoder(body).Decode(&l)

	db, err := config.OpenConec()
	CheckError(err)

	defer db.Close()
	
	sql := `UPDATE livros SET titulo=$2, autor=$3, genero=$4, descricao=$5, preco=$6 WHERE id=$1`

	resul, err := db.Exec(sql, id, l.Titulo, l.Autor, l.Genero, l.Descricao, l.Preco)

	CheckError(err)
	res, _:= resul.RowsAffected()

	return res

}