package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sky0621/cv-admin/generated/swagger"
)

func main() {
	router := swagger.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
