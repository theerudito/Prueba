package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConexionDB() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Nombre de la base de datos
	dbName := os.Getenv("dbHome")

	// Nombre connexión a la base de datos
	dsn := fmt.Sprintf("root:Alfayomega1232*@tcp(127.0.0.1:5588)/%s", dbName)

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
