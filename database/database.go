package database

import (
	"database/sql"
	"fmt"
	"sanctum/util"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser, _                = util.Getenv("DB_USER")
	dbPasswd, _              = util.Getenv("DB_PASSWORD")
	dbName, _                = util.Getenv("DB_NAME")
	connectionStringDatabase = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPasswd, dbName)
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionStringDatabase)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
