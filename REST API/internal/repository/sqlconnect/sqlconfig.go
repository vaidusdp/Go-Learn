package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("Trying to connect SQLServer")
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// connectionString := "sqlserver://sa:vaidus977@localhost:1443?database=" + dbname
	connectionString := fmt.Sprintf("sqlserver://%s:%s@localhost:%s?database=%s", user, password, dbPort, dbname)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to SQLServer")
	return db, nil
}
