package main

import (
	"github.com/jmoiron/sqlx"
)

func main() {
	sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/story")

}
