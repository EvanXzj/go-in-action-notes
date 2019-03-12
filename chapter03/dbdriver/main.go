package main

import (
	"database/sql"
	_ "github.com/EvanXzj/go-in-action-notes/chapter03/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "fakeDataSourceNameString")
}
