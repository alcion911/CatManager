package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    var err error
    // Asegúrate de que la contraseña sea la correcta
    connStr := "user=postgres password=password dbname=animal_db sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }
}