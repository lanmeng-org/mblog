package util

import (
	"path/filepath"
	"html/template"
	"github.com/gin-contrib/multitemplate"
	"strings"
)

func GenRenderer() multitemplate.Renderer {
	layouts := GetCommonTemplateFiles()
	includes := GetTemplateFiles()

	r := multitemplate.NewRenderer()
	for _, include := range includes {
		files := append([]string{include}, layouts...)
		name := strings.Replace(include, BlogConfig.Web.Template.Theme + "/", "", -1)
		name = strings.Replace(name, "\\", "/", -1)
		r.AddFromFilesFuncs(name, GetFuncMap(), files...)
	}

	return r
}

func GetTemplateFiles() []string {
	files := getGlobTemplateFiles([]string{
		BlogConfig.Web.Template.Theme + "/*" + BlogConfig.Web.Template.Suffix,
		BlogConfig.Web.Template.Theme + "/**/*" + BlogConfig.Web.Template.Suffix,
	})

	return files
}

func getGlobTemplateFiles(patterns []string) []string {
	files := []string{}

	for _,p := range patterns {
		f, _ := filepath.Glob(p)
		for _,v := range f {
			name := strings.Replace(v, BlogConfig.Web.Template.Theme + "/", "", -1)
			if string(name[0]) != "_" {
				files = append(files, v)
			}
		}
	}

	return files
}

func GetCommonTemplateFiles() []string {
	files, _ := filepath.Glob(BlogConfig.Web.Template.Theme + "/_**/*" + BlogConfig.Web.Template.Suffix)
	return files
}

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"pageIndexFormat": func(p, i int) int {return (p-1)*10+i+1},
	}
}