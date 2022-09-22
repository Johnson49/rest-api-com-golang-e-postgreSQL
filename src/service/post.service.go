package service

import (
	"encoding/json"
	"io"
	"github.com/Johnson49/rest-api-com-golang-e-postgreSQL/src/config"
	"github.com/Johnson49/rest-api-com-golang-e-postgreSQL/src/models"

	"github.com/google/uuid"
)

func Post(body io.ReadCloser) int64 {
	var l models.Livro
	json.NewDecoder(body).Decode(&l)

	db, err := config.OpenConec()
	CheckError(err)

	defer db.Close()

	sql := `INSERT INTO livros (id, titulo, autor, genero, descricao, preco) values ($1, $2,$3,$4,$5, $6)`

	id := uuid.New().String()
	resul, err := db.Exec(sql, id, l.Titulo, l.Autor, l.Genero, l.Descricao, l.Preco)

	CheckError(err)

	res, err := resul.RowsAffected()
	CheckError(err)

	return res

}
