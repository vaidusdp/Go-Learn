package sqlconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func ConnectDB(dbname string) (*sql.DB, error) {
	fmt.Println("Trying to connect SQLServer")

	connectionString := "sqlserver://sa:vaidus977@localhost:1443?database=" + dbname
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to SQLServer")
	return db, nil
}