//package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "128.199.1.222"
	port     = 5432
	user     = "modulo4"
	password = "modulo4"
	dbname   = "giinwedb"
)

const (
	host1     = "castor.db.elephantsql.com"
	port1     = 5432
	user1     = "imprkjzn"
	password1 = "zjRI3c4nmVoYRRclBfrWiJKCMcuQFa0Z"
	dbname1   = "imprkjzn"
)

func postgreSQLreal() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Conexion exitosa a la BD DeltaIT")

	return db
}

func postgreSQLtest() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host1, port1, user1, password1, dbname1)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Conexion exitosa a la BD de prueba")

	return db
}
