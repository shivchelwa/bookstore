package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"bookstore/models"

	"github.com/ghodss/yaml"
	_ "github.com/lib/pq"
)

func main() {
	var err error

	cfgFileBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(cfgFileBytes))

	config := models.DbStoreConfig{}
	err = yaml.Unmarshal(cfgFileBytes, &config)
	log.Println(config)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.UserName, config.Password, config.DbName)
	log.Printf("trying to connect to database [%s]", psqlInfo)
	models.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")
	http.HandleFunc("/books", getBooks)
	http.ListenAndServe(":3000", nil)
}

// booksIndex sends a HTTP response listing all books.
func getBooks(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
