package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform err %v", err)
		return

	}
	fmt.Fprintf(w, "Post request succesfulll")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name %s\n", name)
	fmt.Fprintf(w, "Address %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 notfaund", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Metodo no soportado ...", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hola!")
}

func main() {
	staticDir := http.Dir("./static")

	staticfileServer := http.StripPrefix("/static/", http.FileServer(staticDir))
	http.Handle("/", staticfileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", formHandler)

	fmt.Printf("Iniciando server en el puerto 8080 ( ochenta ochenta ) \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
