package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func serve(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/tmpl")
	if err != nil {
		log.Println("error parsing template:", err.Error())
		return
	}
	t.Execute(w, "asdf")
}

func data(w http.ResponseWriter, r *http.Request) {
	data := exportData("weather.rrd")
	io.WriteString(w, string(data))
}
