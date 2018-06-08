package main

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"MBlog/util"
)

func main() {


	fmt.Printf("%+v", util.BlogConfig.Web)

}
