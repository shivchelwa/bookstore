package models

import (
	"database/sql"
)

// Create an exported global variable to hold the database connection pool.
var DB *sql.DB

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

type DbStoreConfig struct {
	Host                  string `yaml:"host"`
	Port                  int    `yaml:"port"`
	DbName                string `yaml:"dbname"`
	UserName              string `yaml:"username"`
	Password              string `yaml:"password"`
	MaxOpenConnections    int    `yaml:"dbMaxOpenConnections"`
	MaxIdleConnections    int    `yaml:"dbMaxIdleConnections"`
	ConnectionMaxIdleTime int    `yaml:"dbConnectionMaxIdleTime"`
	QueryTimeout          int    `yaml:"dbQueryTimeout"`
}

// AllBooks returns a slice of all books in the books table.
func AllBooks() ([]Book, error) {
	// Note that we are calling Query() on the global variable.
	rows, err := DB.Query("SELECT * FROM BOOKS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []Book

	for rows.Next() {
		var bk Book

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
