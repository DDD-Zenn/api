package database

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        "db:3306",
        os.Getenv("MYSQL_DATABASE"),
    )

    var err error
    DB, err = sql.Open("mysql", dsn)

    if err != nil {
        return err
    }

	if err := DB.Ping(); err != nil {
		fmt.Println("Failed to connect to database:", err)
		return err
	}

	if err := Migrate(); err != nil {
		return err
	}

	return nil
}