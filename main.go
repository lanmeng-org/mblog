package main

import (
	"MBlog/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"MBlog/util"
	"encoding/hex"
)

func main() {

	pwd, _ := util.EncryptPwd([]byte("123"))

	fmt.Print(pwd)

	user1 := model.User{
		ID: 1,
		Name: "test",
		Password:  hex.EncodeToString(pwd),
		Nickname:  "abc_test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Print(model.SaveUser(&user1))

}
