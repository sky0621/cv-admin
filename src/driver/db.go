package driver

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sky0621/cv-admin/src/adapter/gateway/ent"
)

type CloseDBClientFunc = func()

func MustNewClient() (*ent.Client, CloseDBClientFunc) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	return client, func() { client.Close() }
}
