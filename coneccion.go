package main

import (
	"database/sql"
	"fmt"
)

func ConexionDB() {
	// Datos de conexión
	dsn := "root:Alfayomega1232*@tcp(127.0.0.1:5588)/1010269579_db0000000001"

	// Conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Conexión exitosa a MariaDB")
}
