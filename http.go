package main

import (
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	dataType := vars["type"]
	_ = dataType
	val := vars["val"]
	data := exportData("weather.rrd", val, dataType)
	io.WriteString(w, string(data))
}
