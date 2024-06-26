package tests

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SetupTestDB() *sql.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/referral_db?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Limpar a tabela antes dos testes
	db.Exec("DELETE FROM referrals")

	return db
}
