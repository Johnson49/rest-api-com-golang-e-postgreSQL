package config

import(
	"database/sql"
	"github.com/joho/godotenv"
	"os"
	_ "github.com/lib/pq"
)



func OpenConec() (*sql.DB, error){
	godotenv.Load()
	
	psqlconn := os.Getenv("PSQLCONN")
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

    return db, err

}