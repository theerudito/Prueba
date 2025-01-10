package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() *sql.DB {
	// Datos de conexión
	dsn := "root:Alfayomega1232*@tcp(127.0.0.1:5588)/1010269579_db0000000001"

	// Conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		log.Fatal("Error al verificar la conexión:", err)
	}

	fmt.Println("Conexión exitosa a MariaDB")
	return db
}
