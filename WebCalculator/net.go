package main

import (
	"WebCalculator/calculator"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomePage() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)
	http.HandleFunc("/evaluate", Calculate)
	http.HandleFunc("/", First)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	exp := r.Form.Get("exp")
	res := calculator.CalculatorResult(string(exp))
	t, err := template.ParseFiles("static/result.html")
	if err != nil {
		panic(err)
	}

	data := struct {
		Exp string
		Res string
	}{string(exp), fmt.Sprintf("%f", res)}

	err = t.Execute(w, data)
	// http.ServeFile(w, r, "static/result.html")
}
func First(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
func main() {
	HomePage()
}
