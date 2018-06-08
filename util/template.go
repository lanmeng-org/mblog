package util

import (
	"path/filepath"
	"html/template"
)

func GetTemplateFiles() []string {
	files := []string{}

	for _, path := range BlogConfig.Web.Template.Paths {
		f, _ := filepath.Glob(path + BlogConfig.Web.Template.FileSuffix)
		files = append(files, f...)
	}

	return files
}

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"pageIndexFormat": func(p, i int) int {return (p-1)*10+i+1},
	}
}