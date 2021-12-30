package views 

import (
	"html/template"
	"path/filepath"
	"fmt"
)

var (
	LayoutDir = "views/layouts/"
	TemplateExt = "*.gohtml"
)
// returns slice of strings 
func layoutFiles() []string {
	files, err := filepath.Glob(filepath.Join(LayoutDir, TemplateExt))
	fmt.Println(files)
	if err != nil {
		panic(err)
	}
	return files
}


func NewView(layout string, files ...string) *View {
	//fmt.Println("layoutfiles", layoutFiles())
	files = append(files, layoutFiles()...)
	//fmt.Println(files)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout: layout,
	}
}

type View struct {
	Template *template.Template
	Layout string
}