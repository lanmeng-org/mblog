package util

import (
	"github.com/BurntSushi/toml"
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