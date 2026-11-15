package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Página principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nombre, _ := os.Hostname()
		tmpl.Execute(w, map[string]string{"Pod": nombre})
	})

	// Ruta para comprobar el balanceo (Fase 6)
	http.HandleFunc("/pod", func(w http.ResponseWriter, r *http.Request) {
		nombre, _ := os.Hostname()
		fmt.Fprintf(w, "Respondo desde el pod: %s\n", nombre)
	})

	log.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
