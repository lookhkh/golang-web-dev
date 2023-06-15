package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := struct {
		Name      string
		Motto     string
		CreatedAt time.Time
	}{
		Name:      "Buddha",
		Motto:     "The belief of no beliefs",
		CreatedAt: time.Now(),
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
