package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConnectDB establece la conexión a la base de datos usando los parámetros del archivo .env
func ConnectDB() (*sql.DB, error) {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando archivo .env: %v", err)
	}

	// Crear el DSN (Data Source Name) para la conexión
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")

	// Abrir la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al conectar la base de datos con DSN '%v': %v", dsn, err)
	}

	// Verificar si la base de datos está disponible
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	return db, nil
}
