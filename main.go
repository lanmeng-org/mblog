package main

import (
	_ "github.com/go-sql-driver/mysql"

	"mblog/web"
)

func main() {
	web.Run()
}
