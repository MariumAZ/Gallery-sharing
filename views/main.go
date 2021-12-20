package views 


import (
	"html/template"
	"fmt"
)

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	fmt.Println(files)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}