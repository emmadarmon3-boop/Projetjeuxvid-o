package jeux

import (
	"fmt"
	"net/http"
)

func Server() {
	games := InitGames()
	pseudo := InitUser()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r)
	})
	http.HandleFunc("/connexion", func(w http.ResponseWriter, r *http.Request) {
		ConnexionHandler(w, r, &games, &pseudo)
	})
	http.HandleFunc("/affichage", func(w http.ResponseWriter, r *http.Request) {
		AffichageHandler(w, r)
	})
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server lancé sur http://localhost:8080")
	http.ListenAndServe(":8081", nil)

}
