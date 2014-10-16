package main

import (
	"log"
	"os"
	"text/template"
)

type Student struct {
	// exported field since it begins
	// with a capital letter
	Name string
}

func main() {
	// define an instance
	s := Student{"Marc"}

	// create a new temlate with some name
	tmpl := template.New("test")

	// parse some content and generate a temlate
	tmpl, err := tmpl.Parse("Hello {{.Name}}!")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// merge template 'tmpl' with content os 's'
	err1 := tmpl.Execute(os.Stdout, s)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}
}

