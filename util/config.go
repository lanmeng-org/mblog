package util

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

type configBlog struct {
	DB *configDB `toml:"db"`
	Web *configWeb
}

type configDB struct {
	Type string
	Host string
	Port int
	User string
	Password string
	DBName string
}

func (db *configDB) GenDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		db.User, db.Password, db.Host, db.Port, db.DBName)
}

type configWeb struct {
	Listen string
	Template configWebTemplate
}

type configWebTemplate struct {
	FileSuffix string `toml:"file_suffix"`
	Paths []string
}

var BlogConfig configBlog

func init() {
	_, err := toml.DecodeFile("./config/blog.toml", &BlogConfig)
	if err != nil {
		panic(err)
	}
}