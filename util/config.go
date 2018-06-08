package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type configBlog struct {
	Environment string
	DB          *configDB
	Web         *configWeb
}

type configDB struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Database   string
}

func (db *configDB) GenDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		db.User, db.Password, db.Host, db.Port, db.Database)
}

type configWeb struct {
	Listen   string
	Template configWebTemplate
}

type configWebTemplate struct {
	FileSuffix string
	Paths      []string
}

var BlogConfig configBlog

func init() {
	cFile, err := ioutil.ReadFile("./config/blog.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(cFile, &BlogConfig)
	if err != nil {
		panic(err)
	}
}
