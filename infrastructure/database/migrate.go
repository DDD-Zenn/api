package database

import (
	"fmt"
)

// テーブルが存在しない場合に作成する
func Migrate() error {
	fmt.Println("Running database migration...")  // デバッグログ

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		uid VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	);
	`
	_, err := DB.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return err
	}

	fmt.Println("Table migration complete.")  // デバッグログ
	return nil
}
