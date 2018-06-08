package main

import (
	_ "github.com/go-sql-driver/mysql"

	"MBlog/web"
)

func main() {
	web.Run()
}
