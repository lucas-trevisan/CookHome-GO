package database

import (
	"database/sql"
	"go/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Conn() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.ConnStringDB)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
