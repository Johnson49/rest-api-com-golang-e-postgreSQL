package service

import "github.com/Johnson49/rest-api-com-golang-e-postgreSQL/src/config"

func Delete(id string) int64 {
	db, err := config.OpenConec()
	CheckError(err)

	defer db.Close()
	sql := `DELETE FROM livros WHERE id=$1`

	resul, err := db.Exec(sql, id)
	CheckError(err)

	resultado, err := resul.RowsAffected()
	CheckError(err)
	return resultado
}
