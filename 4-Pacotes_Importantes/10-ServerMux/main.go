package main

import (
	"net/http"
)

/*
	- Ao rotas um http.ListenAndServe o Go vai subir Multi Plexer é um componente onde atachamos as rotas nele.
	- Estamos utilizando o MultiPlexer padrão esse é o mais simples ele é global na aplicação.

	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
*/

func main() {
	// Responsavel em atachar as novas rotas para ter mais controle o recomendavel e criar um mux
	mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World!"))
	//})
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Luiz!"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
