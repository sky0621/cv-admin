package driver

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sky0621/cv-admin/src/ent"
)

type CloseDBClientFunc = func()

func MustNewClient() (*ent.Client, CloseDBClientFunc) {
	// TODO: ent.NewClient 方式にして config で初期設定！
	client, err := ent.Open("sqlite3", "file:db/ent?_fk=1", ent.Debug())
	if err != nil {
		panic(err)
	}
	return client, func() { client.Close() }
}
