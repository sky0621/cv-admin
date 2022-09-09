package main

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sky0621/cv-admin/generated/ent"
)

func main() {
	client, err := ent.Open("sqlite3", "file:db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
