package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/lanmeng-org/mblog/web"
)

func main() {
	web.Run()
}
